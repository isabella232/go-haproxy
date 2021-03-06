package haproxy_test

import (
  "github.com/macb/go-haproxy/haproxy"
  "testing"
)

const (
  testLoadResp   = "# pxname,svname,qcur,qmax,scur,smax,slim,stot,bin,bout,dreq,dresp,ereq,econ,eresp,wretr,wredis,status,weight,act,bck,chkfail,chkdown,lastchg,downtime,qlimit,pid,iid,sid,throttle,lbtot,tracked,type,rate,rate_lim,rate_max,check_status,check_code,check_duration,hrsp_1xx,hrsp_2xx,hrsp_3xx,hrsp_4xx,hrsp_5xx,hrsp_other,hanafail,req_rate,req_rate_max,req_tot,cli_abrt,srv_abrt,comp_in,comp_out,comp_byp,comp_rsp,\nincoming,FRONTEND,,,0,0,2000,0,0,0,0,0,0,,,,,OPEN,,,,,,,,,1,2,0,,,,0,0,0,0,,,,0,0,0,0,0,0,,0,0,0,,,0,0,0,0,\nelastic-ocean,elastic-ocean-1486869-abc,0,0,0,0,32,0,0,0,,0,,0,0,0,0,DOWN,1,1,0,0,1,18,18,,1,3,1,,0,,2,0,,0,L7RSP,,143,0,0,0,0,0,0,0,,,,0,0,,,,,\nelastic-ocean,elastic-ocean-1486869-abc,0,0,0,0,32,0,0,0,,0,,0,0,0,0,DOWN,1,1,0,0,1,18,18,,1,3,1,,0,,2,0,,0,L7RSP,,143,0,0,0,0,0,0,0,,,,0,0,,,,,\nelastic-ocean,elastic-ocean-1486869-abc,0,0,0,0,32,0,0,0,,0,,0,0,0,0,DOWN,1,1,0,0,1,18,18,,1,3,1,,0,,2,0,,0,L7RSP,,143,0,0,0,0,0,0,0,,,,0,0,,,,,\nelastic-ocean,elastic-ocean-1486869-abc,0,0,0,0,32,0,0,0,,0,,0,0,0,0,DOWN,1,1,0,0,1,18,18,,1,3,1,,0,,2,0,,0,L7RSP,,143,0,0,0,0,0,0,0,,,,0,0,,,,,\nelastic-ocean,elastic-ocean-1486869-abc,0,0,0,0,32,0,0,0,,0,,0,0,0,0,DOWN,1,1,0,0,1,18,18,,1,3,1,,0,,2,0,,0,L7RSP,,143,0,0,0,0,0,0,0,,,,0,0,,,,,\nelastic-ocean,BACKEND,0,0,0,0,200,0,0,0,0,0,,0,0,0,0,DOWN,0,0,0,,1,17,17,,1,3,0,,0,,1,0,,0,,,,0,0,0,0,0,0,,,,,0,0,0,0,0,0,"
  testServerName = "elastic-ocean-1486869-abc"
)

func TestGetLoad(t *testing.T) {
  resp := testLoadResp
  tl := ListenForAndRespondWith(t, "show stat\n", resp)

  load, err := haproxy.Haproxy{Socket: haproxy.Socket(tl.Path)}.GetLoad("elastic-ocean")

  if err != nil {
    t.Error(err)
  }
  expectedName := "elastic-ocean-1486869-abc"

  for _, l := range load {
    if l.Name == "BACKEND" {
      continue
    }
    if l.Name != expectedName {
      t.Errorf("Expected %s, got %s.", expectedName, l.Name)
    }
  }
}

func BenchmarkGetLoad(b *testing.B) {
  resp := testLoadResp
  b.StopTimer()
  l, path := ListenForeverAndRespondWith(resp)
  defer l.Close()
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    haproxy.Haproxy{Socket: haproxy.Socket(path)}.GetLoad("elastic-ocean")
  }
}
