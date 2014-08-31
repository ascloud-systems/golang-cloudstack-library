package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type StopVirtualMachineParameter struct {
	Forced NullBool
	Id     ID
}

func (p *StopVirtualMachineParameter) SetForced(b bool) {
	p.Forced.Bool = b
	p.Forced.Valid = true
}
func (p *StopVirtualMachineParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *StopVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Forced.Valid {
		m["forced"] = fmt.Sprint(p.Forced.Bool)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) StopVirtualMachine(p StopVirtualMachineParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("stopVirtualMachine", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	err = json.Unmarshal(v["virtualmachine"], &ret)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	return ret, err
}