// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hearbeat

import (
	"bytes"
	"fmt"
	"github.com/elves-project/agent/src/thrift/apache-thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - ID
//  - IP
//  - Version
//  - Apps
type AgentInfo struct {
	ID      string `thrift:"id,1" json:"id"`
	IP      string `thrift:"ip,2" json:"ip"`
	Version string `thrift:"version,3" json:"version"`
	Apps    string `thrift:"apps,4" json:"apps"`
}

func NewAgentInfo() *AgentInfo {
	return &AgentInfo{}
}

func (p *AgentInfo) GetID() string {
	return p.ID
}

func (p *AgentInfo) GetIP() string {
	return p.IP
}

func (p *AgentInfo) GetVersion() string {
	return p.Version
}

func (p *AgentInfo) GetApps() string {
	return p.Apps
}
func (p *AgentInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AgentInfo) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *AgentInfo) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *AgentInfo) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Version = v
	}
	return nil
}

func (p *AgentInfo) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Apps = v
	}
	return nil
}

func (p *AgentInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("AgentInfo"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AgentInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteString(string(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *AgentInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ip", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ip: ", p), err)
	}
	if err := oprot.WriteString(string(p.IP)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ip (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ip: ", p), err)
	}
	return err
}

func (p *AgentInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("version", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:version: ", p), err)
	}
	if err := oprot.WriteString(string(p.Version)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.version (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:version: ", p), err)
	}
	return err
}

func (p *AgentInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("apps", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:apps: ", p), err)
	}
	if err := oprot.WriteString(string(p.Apps)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.apps (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:apps: ", p), err)
	}
	return err
}

func (p *AgentInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AgentInfo(%+v)", *p)
}
