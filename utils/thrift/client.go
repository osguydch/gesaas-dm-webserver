package thrift

import (
	"compress/zlib"
	"context"
	"crypto/tls"
	"dm/servermanagement"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/spf13/cast"
	"rm/common"
)

var (
	cfg              *thrift.TConfiguration
	transportFactory thrift.TTransportFactory
	protocolFactory  thrift.TProtocolFactory
	ResCode          = map[int32]common.ResDesc{
		1:  common.DmFailedError,
		2:  common.DmInvalidTokenError,
		3:  common.DmNoAuthError,
		4:  common.DmInvalidCreatorIdError,
		5:  common.DmInvalidStatusError,
		6:  common.DmRunIsNotFinishedError,
		7:  common.DmRunIsFinishedError,
		10: common.DmInvalidSampleIdError,
		12: common.DmSampleNameAlreadyExistsError,
		13: common.DmEmptySampleNameError,
		20: common.DmInvalidRunplanIdError,
		21: common.DmRunplanIdAlreadyExistsError,
		30: common.DmInvalidRuntaskIdError,
		40: common.DmInvalidAnalysisidError,
		41: common.DmAnalysisidAlreadyExistsError,
		42: common.DmInvalidRunidError,
		43: common.DmEmptyAnalysisidError,
		44: common.DmEmptyRunidError,
		50: common.DmInvalidDeviceNameError,
		51: common.DmDeviceNameAlreadyExistsError,
		52: common.DmEmptyDeviceNameError,
		60: common.DmInvalidDataTypeError,
		61: common.DmDataTypeAlreadyExistsError,
		62: common.DmDataGroupAlreadyExistsError,
		63: common.DmInvalidDataGroupError,
		64: common.DmDataGroupMappingAlreadyExistsError,
		65: common.DmEmptyGroupNameError,
		70: common.DmInvalidProjectNameError,
		71: common.DmProjectNameAlreadyExistsError,
		80: common.DmInvalidRunDataError,
		81: common.DmRunDataAlreadyExistsError,
		90: common.DmInvalidMacAddressError,
		91: common.DmMacAddressAlreadyExistsError,
		92: common.DmEmptyMacAddressError,
	}
)

func SetupThrift() {
	common.Log.Infof("setup thrift")

	cfg = &thrift.TConfiguration{
		//ConnectTimeout: 60 * time.Second,
		//SocketTimeout:  60 * time.Second,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	switch common.DataManagementConfig.Transport {
	case "http":
		transportFactory = thrift.NewTTransportFactory()
	case "buffered":
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	case "framed":
		transportFactory = thrift.NewTFramedTransportFactoryConf(thrift.NewTTransportFactory(), cfg)
	case "zlib":
		transportFactory = thrift.NewTZlibTransportFactory(zlib.BestCompression)
	default:
		common.Log.Panicf("无效的transport %s;可选值:framed,bufferd,zlib", common.DataManagementConfig.Transport)
	}
	switch common.DataManagementConfig.Protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary":
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
	default:
		common.Log.Panicf("无效的protocol %s;可选值:compact,simplejson,json,binary", common.DataManagementConfig.Protocol)
	}
	tClient, err := newThriftClient(servermanagement.ServerManagementName)
	if err != nil {
		common.Log.Panicf("client dm, create client err：%v", err)
	}

	client := servermanagement.NewServerManagementClient(tClient)
	Response, err := client.GetVersion(context.Background())
	if err != nil {
		common.Log.Panicf("client dm, get version err: %v", err)
	}
	common.Log.Infof("client dm rpc success, dm version: %s", Response.GetData().GetVersion())
}

type client struct {
	serverName string
	transport  thrift.TTransport
}

func newThriftClient(serverName string) (*client, error) {
	hostPort := fmt.Sprintf("%s:%d", common.DataManagementConfig.Ip, common.DataManagementConfig.Port)
	socket := thrift.NewTSocketConf(hostPort, cfg)
	transport, err := transportFactory.GetTransport(socket)
	if err != nil {
		return nil, err
	}
	if err := transport.Open(); err != nil {
		return nil, err
	}

	return &client{
		serverName: serverName,
		transport:  transport,
	}, nil
}

func (p *client) Close() {
	p.transport.Close()
}

func (p *client) Call(ctx context.Context, method string, args, result thrift.TStruct) (thrift.ResponseMeta, error) {
	protocol := protocolFactory.GetProtocol(p.transport)
	serverManagementProtocol := thrift.NewTMultiplexedProtocol(protocol, p.serverName)
	return thrift.NewTStandardClient(serverManagementProtocol, serverManagementProtocol).Call(ctx, method, args, result)
}

type thriftResp interface {
	GetErrorCode() int32
	GetMessage() string
}

func errorHandling(resp thriftResp, err error) error {
	if err != nil {
		return err
	}

	if resp.GetErrorCode() != 0 {
		//dm错误码统一以0x281开头
		code := "0x" + cast.ToString(28100000+resp.GetErrorCode())
		description, ok := ResCode[resp.GetErrorCode()]
		if !ok {
			common.Log.Errorf("errorHandling err: %v", resp)
			description = common.UnknownError
		}
		return &common.ResStatus{Code: code, Description: description}
	}
	return nil
}

var dfCtx = context.Background()
