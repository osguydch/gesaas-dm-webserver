package runmanagement

import (
	"bytes"
	"context"
	"database/sql/driver"
	"dm/rundata"
	"dm/runresult"
	"dm/runtask"
	"dm/sampleinfo"
	"errors"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

var _ = sampleinfo.GoUnusedProtection__
var _ = runtask.GoUnusedProtection__
var _ = runresult.GoUnusedProtection__
var _ = rundata.GoUnusedProtection__

type SampleSheetVersion int64

const (
	SampleSheetVersion_v10 SampleSheetVersion = 1
	SampleSheetVersion_v11 SampleSheetVersion = 2
	SampleSheetVersion_v12 SampleSheetVersion = 3
)

func (p SampleSheetVersion) String() string {
	switch p {
	case SampleSheetVersion_v10:
		return "v10"
	case SampleSheetVersion_v11:
		return "v11"
	case SampleSheetVersion_v12:
		return "v12"
	}
	return "<UNSET>"
}

func SampleSheetVersionFromString(s string) (SampleSheetVersion, error) {
	switch s {
	case "v10":
		return SampleSheetVersion_v10, nil
	case "v11":
		return SampleSheetVersion_v11, nil
	case "v12":
		return SampleSheetVersion_v12, nil
	}
	return SampleSheetVersion(0), fmt.Errorf("not a valid SampleSheetVersion string")
}

func SampleSheetVersionPtr(v SampleSheetVersion) *SampleSheetVersion { return &v }

func (p SampleSheetVersion) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *SampleSheetVersion) UnmarshalText(text []byte) error {
	q, err := SampleSheetVersionFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *SampleSheetVersion) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = SampleSheetVersion(v)
	return nil
}

func (p *SampleSheetVersion) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Attributes:
//   - B
//   - Int16
//   - Int32
//   - Int64
//   - D
//   - S
type GenericField struct {
	B     *bool    `thrift:"b,1" db:"b" json:"b,omitempty"`
	Int16 *int16   `thrift:"int16,2" db:"int16" json:"int16,omitempty"`
	Int32 *int32   `thrift:"int32,3" db:"int32" json:"int32,omitempty"`
	Int64 *int64   `thrift:"int64,4" db:"int64" json:"int64,omitempty"`
	D     *float64 `thrift:"d,5" db:"d" json:"d,omitempty"`
	S     *string  `thrift:"s,6" db:"s" json:"s,omitempty"`
}

func NewGenericField() *GenericField {
	return &GenericField{}
}

var GenericField_B_DEFAULT bool

func (p *GenericField) GetB() bool {
	if !p.IsSetB() {
		return GenericField_B_DEFAULT
	}
	return *p.B
}

var GenericField_Int16_DEFAULT int16

func (p *GenericField) GetInt16() int16 {
	if !p.IsSetInt16() {
		return GenericField_Int16_DEFAULT
	}
	return *p.Int16
}

var GenericField_Int32_DEFAULT int32

func (p *GenericField) GetInt32() int32 {
	if !p.IsSetInt32() {
		return GenericField_Int32_DEFAULT
	}
	return *p.Int32
}

var GenericField_Int64_DEFAULT int64

func (p *GenericField) GetInt64() int64 {
	if !p.IsSetInt64() {
		return GenericField_Int64_DEFAULT
	}
	return *p.Int64
}

var GenericField_D_DEFAULT float64

func (p *GenericField) GetD() float64 {
	if !p.IsSetD() {
		return GenericField_D_DEFAULT
	}
	return *p.D
}

var GenericField_S_DEFAULT string

func (p *GenericField) GetS() string {
	if !p.IsSetS() {
		return GenericField_S_DEFAULT
	}
	return *p.S
}
func (p *GenericField) CountSetFieldsGenericField() int {
	count := 0
	if p.IsSetB() {
		count++
	}
	if p.IsSetInt16() {
		count++
	}
	if p.IsSetInt32() {
		count++
	}
	if p.IsSetInt64() {
		count++
	}
	if p.IsSetD() {
		count++
	}
	if p.IsSetS() {
		count++
	}
	return count

}

func (p *GenericField) IsSetB() bool {
	return p.B != nil
}

func (p *GenericField) IsSetInt16() bool {
	return p.Int16 != nil
}

func (p *GenericField) IsSetInt32() bool {
	return p.Int32 != nil
}

func (p *GenericField) IsSetInt64() bool {
	return p.Int64 != nil
}

func (p *GenericField) IsSetD() bool {
	return p.D != nil
}

func (p *GenericField) IsSetS() bool {
	return p.S != nil
}

func (p *GenericField) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I16 {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.DOUBLE {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField6(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GenericField) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.B = &v
	}
	return nil
}

func (p *GenericField) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Int16 = &v
	}
	return nil
}

func (p *GenericField) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Int32 = &v
	}
	return nil
}

func (p *GenericField) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Int64 = &v
	}
	return nil
}

func (p *GenericField) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(ctx); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.D = &v
	}
	return nil
}

func (p *GenericField) ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.S = &v
	}
	return nil
}

func (p *GenericField) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if c := p.CountSetFieldsGenericField(); c != 1 {
		return fmt.Errorf("%T write union: exactly one field must be set (%d set)", p, c)
	}
	if err := oprot.WriteStructBegin(ctx, "GenericField"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField6(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GenericField) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetB() {
		if err := oprot.WriteFieldBegin(ctx, "b", thrift.BOOL, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:b: ", p), err)
		}
		if err := oprot.WriteBool(ctx, bool(*p.B)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.b (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:b: ", p), err)
		}
	}
	return err
}

func (p *GenericField) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetInt16() {
		if err := oprot.WriteFieldBegin(ctx, "int16", thrift.I16, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:int16: ", p), err)
		}
		if err := oprot.WriteI16(ctx, int16(*p.Int16)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.int16 (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:int16: ", p), err)
		}
	}
	return err
}

func (p *GenericField) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetInt32() {
		if err := oprot.WriteFieldBegin(ctx, "int32", thrift.I32, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:int32: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.Int32)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.int32 (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:int32: ", p), err)
		}
	}
	return err
}

func (p *GenericField) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetInt64() {
		if err := oprot.WriteFieldBegin(ctx, "int64", thrift.I64, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:int64: ", p), err)
		}
		if err := oprot.WriteI64(ctx, int64(*p.Int64)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.int64 (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:int64: ", p), err)
		}
	}
	return err
}

func (p *GenericField) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetD() {
		if err := oprot.WriteFieldBegin(ctx, "d", thrift.DOUBLE, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:d: ", p), err)
		}
		if err := oprot.WriteDouble(ctx, float64(*p.D)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.d (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:d: ", p), err)
		}
	}
	return err
}

func (p *GenericField) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetS() {
		if err := oprot.WriteFieldBegin(ctx, "s", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:s: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.S)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.s (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:s: ", p), err)
		}
	}
	return err
}

func (p *GenericField) Equals(other *GenericField) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.B != other.B {
		if p.B == nil || other.B == nil {
			return false
		}
		if (*p.B) != (*other.B) {
			return false
		}
	}
	if p.Int16 != other.Int16 {
		if p.Int16 == nil || other.Int16 == nil {
			return false
		}
		if (*p.Int16) != (*other.Int16) {
			return false
		}
	}
	if p.Int32 != other.Int32 {
		if p.Int32 == nil || other.Int32 == nil {
			return false
		}
		if (*p.Int32) != (*other.Int32) {
			return false
		}
	}
	if p.Int64 != other.Int64 {
		if p.Int64 == nil || other.Int64 == nil {
			return false
		}
		if (*p.Int64) != (*other.Int64) {
			return false
		}
	}
	if p.D != other.D {
		if p.D == nil || other.D == nil {
			return false
		}
		if (*p.D) != (*other.D) {
			return false
		}
	}
	if p.S != other.S {
		if p.S == nil || other.S == nil {
			return false
		}
		if (*p.S) != (*other.S) {
			return false
		}
	}
	return true
}

func (p *GenericField) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GenericField(%+v)", *p)
}

// Attributes:
//   - RunResult_
//   - RundataList
type RunDataInfo struct {
	RunResult_  *runresult.RunResult_ `thrift:"runResult,1" db:"runResult" json:"runResult"`
	RundataList []*rundata.RunData    `thrift:"rundataList,2" db:"rundataList" json:"rundataList"`
}

func NewRunDataInfo() *RunDataInfo {
	return &RunDataInfo{}
}

var RunDataInfo_RunResult__DEFAULT *runresult.RunResult_

func (p *RunDataInfo) GetRunResult_() *runresult.RunResult_ {
	if !p.IsSetRunResult_() {
		return RunDataInfo_RunResult__DEFAULT
	}
	return p.RunResult_
}

func (p *RunDataInfo) GetRundataList() []*rundata.RunData {
	return p.RundataList
}
func (p *RunDataInfo) IsSetRunResult_() bool {
	return p.RunResult_ != nil
}

func (p *RunDataInfo) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunDataInfo) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.RunResult_ = &runresult.RunResult_{}
	if err := p.RunResult_.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RunResult_), err)
	}
	return nil
}

func (p *RunDataInfo) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*rundata.RunData, 0, size)
	p.RundataList = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &rundata.RunData{}
		if err := _elem0.Read(ctx, iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.RundataList = append(p.RundataList, _elem0)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *RunDataInfo) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RunDataInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunDataInfo) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "runResult", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:runResult: ", p), err)
	}
	if err := p.RunResult_.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RunResult_), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:runResult: ", p), err)
	}
	return err
}

func (p *RunDataInfo) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "rundataList", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:rundataList: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p.RundataList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.RundataList {
		if err := v.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:rundataList: ", p), err)
	}
	return err
}

func (p *RunDataInfo) Equals(other *RunDataInfo) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if !p.RunResult_.Equals(other.RunResult_) {
		return false
	}
	if len(p.RundataList) != len(other.RundataList) {
		return false
	}
	for i, _tgt := range p.RundataList {
		_src1 := other.RundataList[i]
		if !_tgt.Equals(_src1) {
			return false
		}
	}
	return true
}

func (p *RunDataInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunDataInfo(%+v)", *p)
}

// Attributes:
//   - Runtask
//   - RunDataInfoList
//   - TotalSpaceOccupied
//   - ID
type Run struct {
	Runtask            *runtask.RunTask `thrift:"runtask,1" db:"runtask" json:"runtask"`
	RunDataInfoList    []*RunDataInfo   `thrift:"runDataInfoList,2" db:"runDataInfoList" json:"runDataInfoList"`
	TotalSpaceOccupied float64          `thrift:"totalSpaceOccupied,3" db:"totalSpaceOccupied" json:"totalSpaceOccupied"`
	ID                 int32            `thrift:"id,4" db:"id" json:"id"`
}

func NewRun() *Run {
	return &Run{}
}

var Run_Runtask_DEFAULT *runtask.RunTask

func (p *Run) GetRuntask() *runtask.RunTask {
	if !p.IsSetRuntask() {
		return Run_Runtask_DEFAULT
	}
	return p.Runtask
}

func (p *Run) GetRunDataInfoList() []*RunDataInfo {
	return p.RunDataInfoList
}

func (p *Run) GetTotalSpaceOccupied() float64 {
	return p.TotalSpaceOccupied
}

func (p *Run) GetID() int32 {
	return p.ID
}
func (p *Run) IsSetRuntask() bool {
	return p.Runtask != nil
}

func (p *Run) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.DOUBLE {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Run) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Runtask = &runtask.RunTask{}
	if err := p.Runtask.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Runtask), err)
	}
	return nil
}

func (p *Run) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*RunDataInfo, 0, size)
	p.RunDataInfoList = tSlice
	for i := 0; i < size; i++ {
		_elem2 := &RunDataInfo{}
		if err := _elem2.Read(ctx, iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem2), err)
		}
		p.RunDataInfoList = append(p.RunDataInfoList, _elem2)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *Run) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.TotalSpaceOccupied = v
	}
	return nil
}

func (p *Run) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *Run) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Run"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Run) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "runtask", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:runtask: ", p), err)
	}
	if err := p.Runtask.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Runtask), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:runtask: ", p), err)
	}
	return err
}

func (p *Run) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "runDataInfoList", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:runDataInfoList: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p.RunDataInfoList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.RunDataInfoList {
		if err := v.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:runDataInfoList: ", p), err)
	}
	return err
}

func (p *Run) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "totalSpaceOccupied", thrift.DOUBLE, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:totalSpaceOccupied: ", p), err)
	}
	if err := oprot.WriteDouble(ctx, float64(p.TotalSpaceOccupied)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.totalSpaceOccupied (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:totalSpaceOccupied: ", p), err)
	}
	return err
}

func (p *Run) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "id", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:id: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:id: ", p), err)
	}
	return err
}

func (p *Run) Equals(other *Run) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if !p.Runtask.Equals(other.Runtask) {
		return false
	}
	if len(p.RunDataInfoList) != len(other.RunDataInfoList) {
		return false
	}
	for i, _tgt := range p.RunDataInfoList {
		_src3 := other.RunDataInfoList[i]
		if !_tgt.Equals(_src3) {
			return false
		}
	}
	if p.TotalSpaceOccupied != other.TotalSpaceOccupied {
		return false
	}
	if p.ID != other.ID {
		return false
	}
	return true
}

func (p *Run) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Run(%+v)", *p)
}

// Attributes:
//   - ErrorCode
//   - Message
//   - RunList
//   - Analysisid
//   - AnalysisStatus
//   - Samplesheet
//   - TotalCount
//   - CurrentPage
//   - Taskid
type Response struct {
	ErrorCode      int32   `thrift:"errorCode,1,required" db:"errorCode" json:"errorCode"`
	Message        *string `thrift:"message,2" db:"message" json:"message,omitempty"`
	RunList        []*Run  `thrift:"runList,3" db:"runList" json:"runList,omitempty"`
	Analysisid     *string `thrift:"analysisid,4" db:"analysisid" json:"analysisid,omitempty"`
	AnalysisStatus *int32  `thrift:"analysisStatus,5" db:"analysisStatus" json:"analysisStatus,omitempty"`
	Samplesheet    *string `thrift:"samplesheet,6" db:"samplesheet" json:"samplesheet,omitempty"`
	TotalCount     *int32  `thrift:"totalCount,7" db:"totalCount" json:"totalCount,omitempty"`
	CurrentPage    *int32  `thrift:"currentPage,8" db:"currentPage" json:"currentPage,omitempty"`
	Taskid         *string `thrift:"taskid,9" db:"taskid" json:"taskid,omitempty"`
}

func NewResponse() *Response {
	return &Response{}
}

func (p *Response) GetErrorCode() int32 {
	return p.ErrorCode
}

var Response_Message_DEFAULT string

func (p *Response) GetMessage() string {
	if !p.IsSetMessage() {
		return Response_Message_DEFAULT
	}
	return *p.Message
}

var Response_RunList_DEFAULT []*Run

func (p *Response) GetRunList() []*Run {
	return p.RunList
}

var Response_Analysisid_DEFAULT string

func (p *Response) GetAnalysisid() string {
	if !p.IsSetAnalysisid() {
		return Response_Analysisid_DEFAULT
	}
	return *p.Analysisid
}

var Response_AnalysisStatus_DEFAULT int32

func (p *Response) GetAnalysisStatus() int32 {
	if !p.IsSetAnalysisStatus() {
		return Response_AnalysisStatus_DEFAULT
	}
	return *p.AnalysisStatus
}

var Response_Samplesheet_DEFAULT string

func (p *Response) GetSamplesheet() string {
	if !p.IsSetSamplesheet() {
		return Response_Samplesheet_DEFAULT
	}
	return *p.Samplesheet
}

var Response_TotalCount_DEFAULT int32

func (p *Response) GetTotalCount() int32 {
	if !p.IsSetTotalCount() {
		return Response_TotalCount_DEFAULT
	}
	return *p.TotalCount
}

var Response_CurrentPage_DEFAULT int32

func (p *Response) GetCurrentPage() int32 {
	if !p.IsSetCurrentPage() {
		return Response_CurrentPage_DEFAULT
	}
	return *p.CurrentPage
}

var Response_Taskid_DEFAULT string

func (p *Response) GetTaskid() string {
	if !p.IsSetTaskid() {
		return Response_Taskid_DEFAULT
	}
	return *p.Taskid
}
func (p *Response) IsSetMessage() bool {
	return p.Message != nil
}

func (p *Response) IsSetRunList() bool {
	return p.RunList != nil
}

func (p *Response) IsSetAnalysisid() bool {
	return p.Analysisid != nil
}

func (p *Response) IsSetAnalysisStatus() bool {
	return p.AnalysisStatus != nil
}

func (p *Response) IsSetSamplesheet() bool {
	return p.Samplesheet != nil
}

func (p *Response) IsSetTotalCount() bool {
	return p.TotalCount != nil
}

func (p *Response) IsSetCurrentPage() bool {
	return p.CurrentPage != nil
}

func (p *Response) IsSetTaskid() bool {
	return p.Taskid != nil
}

func (p *Response) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetErrorCode bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
				issetErrorCode = true
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField6(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 7:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField7(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 8:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField8(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 9:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField9(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetErrorCode {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ErrorCode is not set"))
	}
	return nil
}

func (p *Response) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ErrorCode = v
	}
	return nil
}

func (p *Response) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Message = &v
	}
	return nil
}

func (p *Response) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Run, 0, size)
	p.RunList = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &Run{}
		if err := _elem4.Read(ctx, iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem4), err)
		}
		p.RunList = append(p.RunList, _elem4)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *Response) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Analysisid = &v
	}
	return nil
}

func (p *Response) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.AnalysisStatus = &v
	}
	return nil
}

func (p *Response) ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Samplesheet = &v
	}
	return nil
}

func (p *Response) ReadField7(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.TotalCount = &v
	}
	return nil
}

func (p *Response) ReadField8(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.CurrentPage = &v
	}
	return nil
}

func (p *Response) ReadField9(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Taskid = &v
	}
	return nil
}

func (p *Response) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Response"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField6(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField7(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField8(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField9(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Response) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "errorCode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:errorCode: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.ErrorCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errorCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:errorCode: ", p), err)
	}
	return err
}

func (p *Response) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetMessage() {
		if err := oprot.WriteFieldBegin(ctx, "message", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:message: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Message)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.message (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:message: ", p), err)
		}
	}
	return err
}

func (p *Response) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetRunList() {
		if err := oprot.WriteFieldBegin(ctx, "runList", thrift.LIST, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:runList: ", p), err)
		}
		if err := oprot.WriteListBegin(ctx, thrift.STRUCT, len(p.RunList)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.RunList {
			if err := v.Write(ctx, oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(ctx); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:runList: ", p), err)
		}
	}
	return err
}

func (p *Response) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetAnalysisid() {
		if err := oprot.WriteFieldBegin(ctx, "analysisid", thrift.STRING, 4); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:analysisid: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Analysisid)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.analysisid (4) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 4:analysisid: ", p), err)
		}
	}
	return err
}

func (p *Response) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetAnalysisStatus() {
		if err := oprot.WriteFieldBegin(ctx, "analysisStatus", thrift.I32, 5); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:analysisStatus: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.AnalysisStatus)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.analysisStatus (5) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 5:analysisStatus: ", p), err)
		}
	}
	return err
}

func (p *Response) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSamplesheet() {
		if err := oprot.WriteFieldBegin(ctx, "samplesheet", thrift.STRING, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:samplesheet: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Samplesheet)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.samplesheet (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:samplesheet: ", p), err)
		}
	}
	return err
}

func (p *Response) writeField7(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetTotalCount() {
		if err := oprot.WriteFieldBegin(ctx, "totalCount", thrift.I32, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:totalCount: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.TotalCount)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.totalCount (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:totalCount: ", p), err)
		}
	}
	return err
}

func (p *Response) writeField8(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetCurrentPage() {
		if err := oprot.WriteFieldBegin(ctx, "currentPage", thrift.I32, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:currentPage: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.CurrentPage)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.currentPage (8) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:currentPage: ", p), err)
		}
	}
	return err
}

func (p *Response) writeField9(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetTaskid() {
		if err := oprot.WriteFieldBegin(ctx, "taskid", thrift.STRING, 9); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:taskid: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Taskid)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.taskid (9) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 9:taskid: ", p), err)
		}
	}
	return err
}

func (p *Response) Equals(other *Response) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.ErrorCode != other.ErrorCode {
		return false
	}
	if p.Message != other.Message {
		if p.Message == nil || other.Message == nil {
			return false
		}
		if (*p.Message) != (*other.Message) {
			return false
		}
	}
	if len(p.RunList) != len(other.RunList) {
		return false
	}
	for i, _tgt := range p.RunList {
		_src5 := other.RunList[i]
		if !_tgt.Equals(_src5) {
			return false
		}
	}
	if p.Analysisid != other.Analysisid {
		if p.Analysisid == nil || other.Analysisid == nil {
			return false
		}
		if (*p.Analysisid) != (*other.Analysisid) {
			return false
		}
	}
	if p.AnalysisStatus != other.AnalysisStatus {
		if p.AnalysisStatus == nil || other.AnalysisStatus == nil {
			return false
		}
		if (*p.AnalysisStatus) != (*other.AnalysisStatus) {
			return false
		}
	}
	if p.Samplesheet != other.Samplesheet {
		if p.Samplesheet == nil || other.Samplesheet == nil {
			return false
		}
		if (*p.Samplesheet) != (*other.Samplesheet) {
			return false
		}
	}
	if p.TotalCount != other.TotalCount {
		if p.TotalCount == nil || other.TotalCount == nil {
			return false
		}
		if (*p.TotalCount) != (*other.TotalCount) {
			return false
		}
	}
	if p.CurrentPage != other.CurrentPage {
		if p.CurrentPage == nil || other.CurrentPage == nil {
			return false
		}
		if (*p.CurrentPage) != (*other.CurrentPage) {
			return false
		}
	}
	if p.Taskid != other.Taskid {
		if p.Taskid == nil || other.Taskid == nil {
			return false
		}
		if (*p.Taskid) != (*other.Taskid) {
			return false
		}
	}
	return true
}

func (p *Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Response(%+v)", *p)
}

type RunManagement interface {
	// Parameters:
	//  - Token
	//  - Runid
	DeleteRun(ctx context.Context, token string, runid string) (_r *Response, _err error)
	// Parameters:
	//  - Token
	//  - Analysisid
	DeleteAnalysis(ctx context.Context, token string, analysisid string) (_r *Response, _err error)
	// Parameters:
	//  - Token
	//  - Runid
	GetRun(ctx context.Context, token string, runid string) (_r *Response, _err error)
	// Parameters:
	//  - Token
	//  - SearchCriteria
	//  - MaxNr
	//  - OrderBy
	GetRunList(ctx context.Context, token string, searchCriteria *runtask.SearchCriteria, maxNr int32, orderBy int32) (_r *Response, _err error)
	// Parameters:
	//  - Token
	//  - Runid
	//  - SampleInfo
	//  - AnalysisParam
	StartOfflineAnalysis(ctx context.Context, token string, runid string, sampleInfo *sampleinfo.SampleInfo, analysisParam map[string]*GenericField) (_r *Response, _err error)
	// Parameters:
	//  - Token
	//  - Analysisid
	StopAnalysis(ctx context.Context, token string, analysisid string) (_r *Response, _err error)
	// Parameters:
	//  - Token
	//  - Analysisid
	GetAnalysisStatus(ctx context.Context, token string, analysisid string) (_r *Response, _err error)
	// Parameters:
	//  - Token
	//  - Analysisid
	//  - Version
	ExportSampleSheetAsCSV(ctx context.Context, token string, analysisid string, version SampleSheetVersion) (_r *Response, _err error)
}

type RunManagementClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewRunManagementClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RunManagementClient {
	return &RunManagementClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewRunManagementClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RunManagementClient {
	return &RunManagementClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewRunManagementClient(c thrift.TClient) *RunManagementClient {
	return &RunManagementClient{
		c: c,
	}
}

func (p *RunManagementClient) Client_() thrift.TClient {
	return p.c
}

func (p *RunManagementClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *RunManagementClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//   - Token
//   - Runid
func (p *RunManagementClient) DeleteRun(ctx context.Context, token string, runid string) (_r *Response, _err error) {
	var _args6 RunManagementDeleteRunArgs
	_args6.Token = token
	_args6.Runid = runid
	var _result8 RunManagementDeleteRunResult
	var _meta7 thrift.ResponseMeta
	_meta7, _err = p.Client_().Call(ctx, "deleteRun", &_args6, &_result8)
	p.SetLastResponseMeta_(_meta7)
	if _err != nil {
		return
	}
	if _ret9 := _result8.GetSuccess(); _ret9 != nil {
		return _ret9, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "deleteRun failed: unknown result")
}

// Parameters:
//   - Token
//   - Analysisid
func (p *RunManagementClient) DeleteAnalysis(ctx context.Context, token string, analysisid string) (_r *Response, _err error) {
	var _args10 RunManagementDeleteAnalysisArgs
	_args10.Token = token
	_args10.Analysisid = analysisid
	var _result12 RunManagementDeleteAnalysisResult
	var _meta11 thrift.ResponseMeta
	_meta11, _err = p.Client_().Call(ctx, "deleteAnalysis", &_args10, &_result12)
	p.SetLastResponseMeta_(_meta11)
	if _err != nil {
		return
	}
	if _ret13 := _result12.GetSuccess(); _ret13 != nil {
		return _ret13, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "deleteAnalysis failed: unknown result")
}

// Parameters:
//   - Token
//   - Runid
func (p *RunManagementClient) GetRun(ctx context.Context, token string, runid string) (_r *Response, _err error) {
	var _args14 RunManagementGetRunArgs
	_args14.Token = token
	_args14.Runid = runid
	var _result16 RunManagementGetRunResult
	var _meta15 thrift.ResponseMeta
	_meta15, _err = p.Client_().Call(ctx, "getRun", &_args14, &_result16)
	p.SetLastResponseMeta_(_meta15)
	if _err != nil {
		return
	}
	if _ret17 := _result16.GetSuccess(); _ret17 != nil {
		return _ret17, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "getRun failed: unknown result")
}

// Parameters:
//   - Token
//   - SearchCriteria
//   - MaxNr
//   - OrderBy
func (p *RunManagementClient) GetRunList(ctx context.Context, token string, searchCriteria *runtask.SearchCriteria, maxNr int32, orderBy int32) (_r *Response, _err error) {
	var _args18 RunManagementGetRunListArgs
	_args18.Token = token
	_args18.SearchCriteria = searchCriteria
	_args18.MaxNr = maxNr
	_args18.OrderBy = orderBy
	var _result20 RunManagementGetRunListResult
	var _meta19 thrift.ResponseMeta
	_meta19, _err = p.Client_().Call(ctx, "getRunList", &_args18, &_result20)
	p.SetLastResponseMeta_(_meta19)
	if _err != nil {
		return
	}
	if _ret21 := _result20.GetSuccess(); _ret21 != nil {
		return _ret21, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "getRunList failed: unknown result")
}

// Parameters:
//   - Token
//   - Runid
//   - SampleInfo
//   - AnalysisParam
func (p *RunManagementClient) StartOfflineAnalysis(ctx context.Context, token string, runid string, sampleInfo *sampleinfo.SampleInfo, analysisParam map[string]*GenericField) (_r *Response, _err error) {
	var _args22 RunManagementStartOfflineAnalysisArgs
	_args22.Token = token
	_args22.Runid = runid
	_args22.SampleInfo = sampleInfo
	_args22.AnalysisParam = analysisParam
	var _result24 RunManagementStartOfflineAnalysisResult
	var _meta23 thrift.ResponseMeta
	_meta23, _err = p.Client_().Call(ctx, "startOfflineAnalysis", &_args22, &_result24)
	p.SetLastResponseMeta_(_meta23)
	if _err != nil {
		return
	}
	if _ret25 := _result24.GetSuccess(); _ret25 != nil {
		return _ret25, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "startOfflineAnalysis failed: unknown result")
}

// Parameters:
//   - Token
//   - Analysisid
func (p *RunManagementClient) StopAnalysis(ctx context.Context, token string, analysisid string) (_r *Response, _err error) {
	var _args26 RunManagementStopAnalysisArgs
	_args26.Token = token
	_args26.Analysisid = analysisid
	var _result28 RunManagementStopAnalysisResult
	var _meta27 thrift.ResponseMeta
	_meta27, _err = p.Client_().Call(ctx, "stopAnalysis", &_args26, &_result28)
	p.SetLastResponseMeta_(_meta27)
	if _err != nil {
		return
	}
	if _ret29 := _result28.GetSuccess(); _ret29 != nil {
		return _ret29, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "stopAnalysis failed: unknown result")
}

// Parameters:
//   - Token
//   - Analysisid
func (p *RunManagementClient) GetAnalysisStatus(ctx context.Context, token string, analysisid string) (_r *Response, _err error) {
	var _args30 RunManagementGetAnalysisStatusArgs
	_args30.Token = token
	_args30.Analysisid = analysisid
	var _result32 RunManagementGetAnalysisStatusResult
	var _meta31 thrift.ResponseMeta
	_meta31, _err = p.Client_().Call(ctx, "getAnalysisStatus", &_args30, &_result32)
	p.SetLastResponseMeta_(_meta31)
	if _err != nil {
		return
	}
	if _ret33 := _result32.GetSuccess(); _ret33 != nil {
		return _ret33, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "getAnalysisStatus failed: unknown result")
}

// Parameters:
//   - Token
//   - Analysisid
//   - Version
func (p *RunManagementClient) ExportSampleSheetAsCSV(ctx context.Context, token string, analysisid string, version SampleSheetVersion) (_r *Response, _err error) {
	var _args34 RunManagementExportSampleSheetAsCSVArgs
	_args34.Token = token
	_args34.Analysisid = analysisid
	_args34.Version = version
	var _result36 RunManagementExportSampleSheetAsCSVResult
	var _meta35 thrift.ResponseMeta
	_meta35, _err = p.Client_().Call(ctx, "exportSampleSheetAsCSV", &_args34, &_result36)
	p.SetLastResponseMeta_(_meta35)
	if _err != nil {
		return
	}
	if _ret37 := _result36.GetSuccess(); _ret37 != nil {
		return _ret37, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "exportSampleSheetAsCSV failed: unknown result")
}

type RunManagementProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      RunManagement
}

func (p *RunManagementProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *RunManagementProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *RunManagementProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewRunManagementProcessor(handler RunManagement) *RunManagementProcessor {

	self38 := &RunManagementProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self38.processorMap["deleteRun"] = &runManagementProcessorDeleteRun{handler: handler}
	self38.processorMap["deleteAnalysis"] = &runManagementProcessorDeleteAnalysis{handler: handler}
	self38.processorMap["getRun"] = &runManagementProcessorGetRun{handler: handler}
	self38.processorMap["getRunList"] = &runManagementProcessorGetRunList{handler: handler}
	self38.processorMap["startOfflineAnalysis"] = &runManagementProcessorStartOfflineAnalysis{handler: handler}
	self38.processorMap["stopAnalysis"] = &runManagementProcessorStopAnalysis{handler: handler}
	self38.processorMap["getAnalysisStatus"] = &runManagementProcessorGetAnalysisStatus{handler: handler}
	self38.processorMap["exportSampleSheetAsCSV"] = &runManagementProcessorExportSampleSheetAsCSV{handler: handler}
	return self38
}

func (p *RunManagementProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x39 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x39.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x39

}

type runManagementProcessorDeleteRun struct {
	handler RunManagement
}

func (p *runManagementProcessorDeleteRun) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementDeleteRunArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "deleteRun", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementDeleteRunResult{}
	var retval *Response
	if retval, err2 = p.handler.DeleteRun(ctx, args.Token, args.Runid); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteRun: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "deleteRun", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "deleteRun", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type runManagementProcessorDeleteAnalysis struct {
	handler RunManagement
}

func (p *runManagementProcessorDeleteAnalysis) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementDeleteAnalysisArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "deleteAnalysis", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementDeleteAnalysisResult{}
	var retval *Response
	if retval, err2 = p.handler.DeleteAnalysis(ctx, args.Token, args.Analysisid); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing deleteAnalysis: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "deleteAnalysis", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "deleteAnalysis", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type runManagementProcessorGetRun struct {
	handler RunManagement
}

func (p *runManagementProcessorGetRun) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementGetRunArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "getRun", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementGetRunResult{}
	var retval *Response
	if retval, err2 = p.handler.GetRun(ctx, args.Token, args.Runid); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRun: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "getRun", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "getRun", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type runManagementProcessorGetRunList struct {
	handler RunManagement
}

func (p *runManagementProcessorGetRunList) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementGetRunListArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "getRunList", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementGetRunListResult{}
	var retval *Response
	if retval, err2 = p.handler.GetRunList(ctx, args.Token, args.SearchCriteria, args.MaxNr, args.OrderBy); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getRunList: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "getRunList", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "getRunList", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type runManagementProcessorStartOfflineAnalysis struct {
	handler RunManagement
}

func (p *runManagementProcessorStartOfflineAnalysis) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementStartOfflineAnalysisArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "startOfflineAnalysis", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementStartOfflineAnalysisResult{}
	var retval *Response
	if retval, err2 = p.handler.StartOfflineAnalysis(ctx, args.Token, args.Runid, args.SampleInfo, args.AnalysisParam); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing startOfflineAnalysis: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "startOfflineAnalysis", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "startOfflineAnalysis", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type runManagementProcessorStopAnalysis struct {
	handler RunManagement
}

func (p *runManagementProcessorStopAnalysis) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementStopAnalysisArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "stopAnalysis", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementStopAnalysisResult{}
	var retval *Response
	if retval, err2 = p.handler.StopAnalysis(ctx, args.Token, args.Analysisid); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing stopAnalysis: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "stopAnalysis", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "stopAnalysis", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type runManagementProcessorGetAnalysisStatus struct {
	handler RunManagement
}

func (p *runManagementProcessorGetAnalysisStatus) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementGetAnalysisStatusArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "getAnalysisStatus", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementGetAnalysisStatusResult{}
	var retval *Response
	if retval, err2 = p.handler.GetAnalysisStatus(ctx, args.Token, args.Analysisid); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getAnalysisStatus: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "getAnalysisStatus", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "getAnalysisStatus", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type runManagementProcessorExportSampleSheetAsCSV struct {
	handler RunManagement
}

func (p *runManagementProcessorExportSampleSheetAsCSV) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RunManagementExportSampleSheetAsCSVArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "exportSampleSheetAsCSV", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := RunManagementExportSampleSheetAsCSVResult{}
	var retval *Response
	if retval, err2 = p.handler.ExportSampleSheetAsCSV(ctx, args.Token, args.Analysisid, args.Version); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing exportSampleSheetAsCSV: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "exportSampleSheetAsCSV", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "exportSampleSheetAsCSV", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//   - Token
//   - Runid
type RunManagementDeleteRunArgs struct {
	Token string `thrift:"token,1" db:"token" json:"token"`
	Runid string `thrift:"runid,2" db:"runid" json:"runid"`
}

func NewRunManagementDeleteRunArgs() *RunManagementDeleteRunArgs {
	return &RunManagementDeleteRunArgs{}
}

func (p *RunManagementDeleteRunArgs) GetToken() string {
	return p.Token
}

func (p *RunManagementDeleteRunArgs) GetRunid() string {
	return p.Runid
}
func (p *RunManagementDeleteRunArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementDeleteRunArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementDeleteRunArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Runid = v
	}
	return nil
}

func (p *RunManagementDeleteRunArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "deleteRun_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementDeleteRunArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementDeleteRunArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "runid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:runid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Runid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.runid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:runid: ", p), err)
	}
	return err
}

func (p *RunManagementDeleteRunArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementDeleteRunArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementDeleteRunResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementDeleteRunResult() *RunManagementDeleteRunResult {
	return &RunManagementDeleteRunResult{}
}

var RunManagementDeleteRunResult_Success_DEFAULT *Response

func (p *RunManagementDeleteRunResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementDeleteRunResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementDeleteRunResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementDeleteRunResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementDeleteRunResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementDeleteRunResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "deleteRun_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementDeleteRunResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementDeleteRunResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementDeleteRunResult(%+v)", *p)
}

// Attributes:
//   - Token
//   - Analysisid
type RunManagementDeleteAnalysisArgs struct {
	Token      string `thrift:"token,1" db:"token" json:"token"`
	Analysisid string `thrift:"analysisid,2" db:"analysisid" json:"analysisid"`
}

func NewRunManagementDeleteAnalysisArgs() *RunManagementDeleteAnalysisArgs {
	return &RunManagementDeleteAnalysisArgs{}
}

func (p *RunManagementDeleteAnalysisArgs) GetToken() string {
	return p.Token
}

func (p *RunManagementDeleteAnalysisArgs) GetAnalysisid() string {
	return p.Analysisid
}
func (p *RunManagementDeleteAnalysisArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementDeleteAnalysisArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementDeleteAnalysisArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Analysisid = v
	}
	return nil
}

func (p *RunManagementDeleteAnalysisArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "deleteAnalysis_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementDeleteAnalysisArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementDeleteAnalysisArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "analysisid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:analysisid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Analysisid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.analysisid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:analysisid: ", p), err)
	}
	return err
}

func (p *RunManagementDeleteAnalysisArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementDeleteAnalysisArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementDeleteAnalysisResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementDeleteAnalysisResult() *RunManagementDeleteAnalysisResult {
	return &RunManagementDeleteAnalysisResult{}
}

var RunManagementDeleteAnalysisResult_Success_DEFAULT *Response

func (p *RunManagementDeleteAnalysisResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementDeleteAnalysisResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementDeleteAnalysisResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementDeleteAnalysisResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementDeleteAnalysisResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementDeleteAnalysisResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "deleteAnalysis_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementDeleteAnalysisResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementDeleteAnalysisResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementDeleteAnalysisResult(%+v)", *p)
}

// Attributes:
//   - Token
//   - Runid
type RunManagementGetRunArgs struct {
	Token string `thrift:"token,1" db:"token" json:"token"`
	Runid string `thrift:"runid,2" db:"runid" json:"runid"`
}

func NewRunManagementGetRunArgs() *RunManagementGetRunArgs {
	return &RunManagementGetRunArgs{}
}

func (p *RunManagementGetRunArgs) GetToken() string {
	return p.Token
}

func (p *RunManagementGetRunArgs) GetRunid() string {
	return p.Runid
}
func (p *RunManagementGetRunArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementGetRunArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementGetRunArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Runid = v
	}
	return nil
}

func (p *RunManagementGetRunArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getRun_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementGetRunArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementGetRunArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "runid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:runid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Runid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.runid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:runid: ", p), err)
	}
	return err
}

func (p *RunManagementGetRunArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementGetRunArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementGetRunResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementGetRunResult() *RunManagementGetRunResult {
	return &RunManagementGetRunResult{}
}

var RunManagementGetRunResult_Success_DEFAULT *Response

func (p *RunManagementGetRunResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementGetRunResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementGetRunResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementGetRunResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementGetRunResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementGetRunResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getRun_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementGetRunResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementGetRunResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementGetRunResult(%+v)", *p)
}

// Attributes:
//   - Token
//   - SearchCriteria
//   - MaxNr
//   - OrderBy
type RunManagementGetRunListArgs struct {
	Token          string                  `thrift:"token,1" db:"token" json:"token"`
	SearchCriteria *runtask.SearchCriteria `thrift:"searchCriteria,2" db:"searchCriteria" json:"searchCriteria"`
	MaxNr          int32                   `thrift:"maxNr,3" db:"maxNr" json:"maxNr"`
	OrderBy        int32                   `thrift:"orderBy,4" db:"orderBy" json:"orderBy"`
}

func NewRunManagementGetRunListArgs() *RunManagementGetRunListArgs {
	return &RunManagementGetRunListArgs{}
}

func (p *RunManagementGetRunListArgs) GetToken() string {
	return p.Token
}

var RunManagementGetRunListArgs_SearchCriteria_DEFAULT *runtask.SearchCriteria

func (p *RunManagementGetRunListArgs) GetSearchCriteria() *runtask.SearchCriteria {
	if !p.IsSetSearchCriteria() {
		return RunManagementGetRunListArgs_SearchCriteria_DEFAULT
	}
	return p.SearchCriteria
}

func (p *RunManagementGetRunListArgs) GetMaxNr() int32 {
	return p.MaxNr
}

func (p *RunManagementGetRunListArgs) GetOrderBy() int32 {
	return p.OrderBy
}
func (p *RunManagementGetRunListArgs) IsSetSearchCriteria() bool {
	return p.SearchCriteria != nil
}

func (p *RunManagementGetRunListArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementGetRunListArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementGetRunListArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.SearchCriteria = &runtask.SearchCriteria{}
	if err := p.SearchCriteria.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.SearchCriteria), err)
	}
	return nil
}

func (p *RunManagementGetRunListArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.MaxNr = v
	}
	return nil
}

func (p *RunManagementGetRunListArgs) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.OrderBy = v
	}
	return nil
}

func (p *RunManagementGetRunListArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getRunList_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementGetRunListArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementGetRunListArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "searchCriteria", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:searchCriteria: ", p), err)
	}
	if err := p.SearchCriteria.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.SearchCriteria), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:searchCriteria: ", p), err)
	}
	return err
}

func (p *RunManagementGetRunListArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "maxNr", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:maxNr: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.MaxNr)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.maxNr (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:maxNr: ", p), err)
	}
	return err
}

func (p *RunManagementGetRunListArgs) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "orderBy", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:orderBy: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.OrderBy)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.orderBy (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:orderBy: ", p), err)
	}
	return err
}

func (p *RunManagementGetRunListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementGetRunListArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementGetRunListResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementGetRunListResult() *RunManagementGetRunListResult {
	return &RunManagementGetRunListResult{}
}

var RunManagementGetRunListResult_Success_DEFAULT *Response

func (p *RunManagementGetRunListResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementGetRunListResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementGetRunListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementGetRunListResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementGetRunListResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementGetRunListResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getRunList_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementGetRunListResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementGetRunListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementGetRunListResult(%+v)", *p)
}

// Attributes:
//   - Token
//   - Runid
//   - SampleInfo
//   - AnalysisParam
type RunManagementStartOfflineAnalysisArgs struct {
	Token         string                   `thrift:"token,1" db:"token" json:"token"`
	Runid         string                   `thrift:"runid,2" db:"runid" json:"runid"`
	SampleInfo    *sampleinfo.SampleInfo   `thrift:"sampleInfo,3" db:"sampleInfo" json:"sampleInfo"`
	AnalysisParam map[string]*GenericField `thrift:"analysisParam,4" db:"analysisParam" json:"analysisParam"`
}

func NewRunManagementStartOfflineAnalysisArgs() *RunManagementStartOfflineAnalysisArgs {
	return &RunManagementStartOfflineAnalysisArgs{}
}

func (p *RunManagementStartOfflineAnalysisArgs) GetToken() string {
	return p.Token
}

func (p *RunManagementStartOfflineAnalysisArgs) GetRunid() string {
	return p.Runid
}

var RunManagementStartOfflineAnalysisArgs_SampleInfo_DEFAULT *sampleinfo.SampleInfo

func (p *RunManagementStartOfflineAnalysisArgs) GetSampleInfo() *sampleinfo.SampleInfo {
	if !p.IsSetSampleInfo() {
		return RunManagementStartOfflineAnalysisArgs_SampleInfo_DEFAULT
	}
	return p.SampleInfo
}

func (p *RunManagementStartOfflineAnalysisArgs) GetAnalysisParam() map[string]*GenericField {
	return p.AnalysisParam
}
func (p *RunManagementStartOfflineAnalysisArgs) IsSetSampleInfo() bool {
	return p.SampleInfo != nil
}

func (p *RunManagementStartOfflineAnalysisArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.MAP {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Runid = v
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.SampleInfo = &sampleinfo.SampleInfo{}
	if err := p.SampleInfo.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.SampleInfo), err)
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisArgs) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]*GenericField, size)
	p.AnalysisParam = tMap
	for i := 0; i < size; i++ {
		var _key40 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key40 = v
		}
		_val41 := &GenericField{}
		if err := _val41.Read(ctx, iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _val41), err)
		}
		p.AnalysisParam[_key40] = _val41
	}
	if err := iprot.ReadMapEnd(ctx); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "startOfflineAnalysis_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementStartOfflineAnalysisArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "runid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:runid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Runid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.runid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:runid: ", p), err)
	}
	return err
}

func (p *RunManagementStartOfflineAnalysisArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sampleInfo", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:sampleInfo: ", p), err)
	}
	if err := p.SampleInfo.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.SampleInfo), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:sampleInfo: ", p), err)
	}
	return err
}

func (p *RunManagementStartOfflineAnalysisArgs) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "analysisParam", thrift.MAP, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:analysisParam: ", p), err)
	}
	if err := oprot.WriteMapBegin(ctx, thrift.STRING, thrift.STRUCT, len(p.AnalysisParam)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.AnalysisParam {
		if err := oprot.WriteString(ctx, string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := v.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteMapEnd(ctx); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:analysisParam: ", p), err)
	}
	return err
}

func (p *RunManagementStartOfflineAnalysisArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementStartOfflineAnalysisArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementStartOfflineAnalysisResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementStartOfflineAnalysisResult() *RunManagementStartOfflineAnalysisResult {
	return &RunManagementStartOfflineAnalysisResult{}
}

var RunManagementStartOfflineAnalysisResult_Success_DEFAULT *Response

func (p *RunManagementStartOfflineAnalysisResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementStartOfflineAnalysisResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementStartOfflineAnalysisResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementStartOfflineAnalysisResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "startOfflineAnalysis_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementStartOfflineAnalysisResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementStartOfflineAnalysisResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementStartOfflineAnalysisResult(%+v)", *p)
}

// Attributes:
//   - Token
//   - Analysisid
type RunManagementStopAnalysisArgs struct {
	Token      string `thrift:"token,1" db:"token" json:"token"`
	Analysisid string `thrift:"analysisid,2" db:"analysisid" json:"analysisid"`
}

func NewRunManagementStopAnalysisArgs() *RunManagementStopAnalysisArgs {
	return &RunManagementStopAnalysisArgs{}
}

func (p *RunManagementStopAnalysisArgs) GetToken() string {
	return p.Token
}

func (p *RunManagementStopAnalysisArgs) GetAnalysisid() string {
	return p.Analysisid
}
func (p *RunManagementStopAnalysisArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementStopAnalysisArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementStopAnalysisArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Analysisid = v
	}
	return nil
}

func (p *RunManagementStopAnalysisArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "stopAnalysis_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementStopAnalysisArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementStopAnalysisArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "analysisid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:analysisid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Analysisid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.analysisid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:analysisid: ", p), err)
	}
	return err
}

func (p *RunManagementStopAnalysisArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementStopAnalysisArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementStopAnalysisResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementStopAnalysisResult() *RunManagementStopAnalysisResult {
	return &RunManagementStopAnalysisResult{}
}

var RunManagementStopAnalysisResult_Success_DEFAULT *Response

func (p *RunManagementStopAnalysisResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementStopAnalysisResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementStopAnalysisResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementStopAnalysisResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementStopAnalysisResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementStopAnalysisResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "stopAnalysis_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementStopAnalysisResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementStopAnalysisResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementStopAnalysisResult(%+v)", *p)
}

// Attributes:
//   - Token
//   - Analysisid
type RunManagementGetAnalysisStatusArgs struct {
	Token      string `thrift:"token,1" db:"token" json:"token"`
	Analysisid string `thrift:"analysisid,2" db:"analysisid" json:"analysisid"`
}

func NewRunManagementGetAnalysisStatusArgs() *RunManagementGetAnalysisStatusArgs {
	return &RunManagementGetAnalysisStatusArgs{}
}

func (p *RunManagementGetAnalysisStatusArgs) GetToken() string {
	return p.Token
}

func (p *RunManagementGetAnalysisStatusArgs) GetAnalysisid() string {
	return p.Analysisid
}
func (p *RunManagementGetAnalysisStatusArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementGetAnalysisStatusArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementGetAnalysisStatusArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Analysisid = v
	}
	return nil
}

func (p *RunManagementGetAnalysisStatusArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getAnalysisStatus_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementGetAnalysisStatusArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementGetAnalysisStatusArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "analysisid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:analysisid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Analysisid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.analysisid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:analysisid: ", p), err)
	}
	return err
}

func (p *RunManagementGetAnalysisStatusArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementGetAnalysisStatusArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementGetAnalysisStatusResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementGetAnalysisStatusResult() *RunManagementGetAnalysisStatusResult {
	return &RunManagementGetAnalysisStatusResult{}
}

var RunManagementGetAnalysisStatusResult_Success_DEFAULT *Response

func (p *RunManagementGetAnalysisStatusResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementGetAnalysisStatusResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementGetAnalysisStatusResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementGetAnalysisStatusResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementGetAnalysisStatusResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementGetAnalysisStatusResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getAnalysisStatus_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementGetAnalysisStatusResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementGetAnalysisStatusResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementGetAnalysisStatusResult(%+v)", *p)
}

// Attributes:
//   - Token
//   - Analysisid
//   - Version
type RunManagementExportSampleSheetAsCSVArgs struct {
	Token      string             `thrift:"token,1" db:"token" json:"token"`
	Analysisid string             `thrift:"analysisid,2" db:"analysisid" json:"analysisid"`
	Version    SampleSheetVersion `thrift:"version,3" db:"version" json:"version"`
}

func NewRunManagementExportSampleSheetAsCSVArgs() *RunManagementExportSampleSheetAsCSVArgs {
	return &RunManagementExportSampleSheetAsCSVArgs{}
}

func (p *RunManagementExportSampleSheetAsCSVArgs) GetToken() string {
	return p.Token
}

func (p *RunManagementExportSampleSheetAsCSVArgs) GetAnalysisid() string {
	return p.Analysisid
}

func (p *RunManagementExportSampleSheetAsCSVArgs) GetVersion() SampleSheetVersion {
	return p.Version
}
func (p *RunManagementExportSampleSheetAsCSVArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Analysisid = v
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := SampleSheetVersion(v)
		p.Version = temp
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "exportSampleSheetAsCSV_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *RunManagementExportSampleSheetAsCSVArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "analysisid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:analysisid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Analysisid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.analysisid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:analysisid: ", p), err)
	}
	return err
}

func (p *RunManagementExportSampleSheetAsCSVArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "version", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:version: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Version)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.version (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:version: ", p), err)
	}
	return err
}

func (p *RunManagementExportSampleSheetAsCSVArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementExportSampleSheetAsCSVArgs(%+v)", *p)
}

// Attributes:
//   - Success
type RunManagementExportSampleSheetAsCSVResult struct {
	Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewRunManagementExportSampleSheetAsCSVResult() *RunManagementExportSampleSheetAsCSVResult {
	return &RunManagementExportSampleSheetAsCSVResult{}
}

var RunManagementExportSampleSheetAsCSVResult_Success_DEFAULT *Response

func (p *RunManagementExportSampleSheetAsCSVResult) GetSuccess() *Response {
	if !p.IsSetSuccess() {
		return RunManagementExportSampleSheetAsCSVResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RunManagementExportSampleSheetAsCSVResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RunManagementExportSampleSheetAsCSVResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "exportSampleSheetAsCSV_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RunManagementExportSampleSheetAsCSVResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RunManagementExportSampleSheetAsCSVResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RunManagementExportSampleSheetAsCSVResult(%+v)", *p)
}
