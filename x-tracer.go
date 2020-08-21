
package main

import (
        "context"
        "fmt"
        "strings"
        "github.com/docker/docker/api/types"
        "github.com/docker/docker/client"
        pp "github/Sheenam3/techday/parse/probeparser"
)

func main() {
        ctx := context.Background()
        cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
        if err != nil {
                panic(err)
        }

        containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
        if err != nil {
                panic(err)
        }
        fmt.Println("-------------------Choose Container-----------------------------")
        for index, container := range containers {
                out := strings.TrimLeft(container.Names[0],"/")
                fmt.Println(index, ":", out)
        }

        fmt.Print("Choose Container: ")

        var con int
        _, err = fmt.Scanf("%d", &con)

        if err != nil {
                fmt.Println(err)
        }


        fmt.Println("---------------------------------------------")
        fmt.Println("The Container you chose is:", strings.TrimLeft(containers[con].Names[0],"/"))
        fmt.Println("Container Id is:", containers[con].ID)

        //Tools/ Probes
        pn := []string {"tcptracer", "tcpconnect", "tcpaccept", "tcplife", "execsnoop", "biosnoop", "cachestat", "All Probes"}

        topResult, err := cli.ContainerTop(context.Background(), containers[con].ID, []string{"o","pid"})
        if err != nil {
                panic(err)
        }
        fmt.Println(topResult.Processes)

        //Run Probes
        fmt.Println("---------------------------------------------")
        fmt.Println("-------------------Choose Probe-----------------------------")

        for index, proben :=  range pn{
                fmt.Println(index, ":", proben)

        }
        fmt.Print("Choose Probe: ")
        var probe int
        _, err = fmt.Scanf("%d", &probe)

        if err != nil {
                fmt.Println(err)
        }



        fmt.Println("The Probe you chose is:", pn[probe])


        switch pn[probe] {

        case "tcptracer":
                logtcptracer := make(chan pp.Log, 1)
                        go pp.RunTcptracer(pn[probe], logtcptracer, topResult[0][0])
                        go func() {

                                for val := range logtcptracer {

                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s |IP->%s | SADDR:%s | DADDR:%s | SPORT:%s | DPORT:%s \n",pn[probe],val.fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8],val.Fulllog[9])


                                }

                        }()

        case "tcpconnect":
                logtcpconnect := make(chan pp.Log, 1)
                        go pp.RunTcpconnect(pn[probe], logtcpconnect, topResult[0][0])
                        go func() {

                                for val := range logtcpconnect {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8])
                                }

                        }()


        case "tcpaccept":
                logtcpaccept := make(chan pp.Log, 1)
                        go pp.RunTcpaccept(pn[probe], logtcpaccept, topResult[0][0])
                        go func() {

                                for val := range logtcpaccept {

                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | RADDR:%s | RPORT:%s | LADDR:%s | LPORT:%s \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8],val.Fulllog[9])


                                }

                        }()

        case "tcplife":
                logtcplife := make(chan pp.Log, 1)
                        go pp.RunTcplife(pn[probe], logtcplife, topResult[0][0])
                        go func() {

                                for val := range logtcplife {

                                        fmt.Printf("{Probe:%s |Sys_Time: %s |PID:%s | PNAME:%s | LADDRR:%s | LPORT:%s | RADDR:%s | RPORT:%s | TX_KB:%s | RX_KB:%s | MS: %s \n",pn[probe],val.Fulllog[0],val.Fulllog[2],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8],val.Fulllog[9],val.Fulllog[10])

                                }

                        }()
        case "execsnoop":
                logexecsnoop := make(chan pp.Log, 1)
                        go pp.RunExecsnoop(pn[probe], logexecsnoop, topResult[0][0])
                        go func() {

                                for val := range logexecsnoop {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s | T:%s | PNAME: %s | PID:%s | PPID:%s | RET:%s | ARGS:%s \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7])

                                }

                        }()
        case "biosnoop":
                logbiosnoop := make(chan pp.Log, 1)
                        go pp.RunBiosnoop(pn[probe], logbiosnoop, topResult[0][0])
                        go func() {

                                for val := range logbiosnoop {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s |PNAME: %s | PID:%s | DISK:%s | R/W:%s | SECTOR:%s |BYTES: %s | Lat(ms): %s | \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[2],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[9])


                                }

                        }()
        case "cachestat":
                logcachetop := make(chan pp.Log, 1)
                        go pp.RunCachetop(pn[probe], logcachetop, topResult[0][0])
                        go func() {

                                for val := range logcachetop {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s | PID:%s | UID:%s | CMD:%s | HITS:%s | MISS:%s | DIRTIES: %s| READ_HIT%:%s | W_HIT%: %s | \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[2],val.Fulllog[3],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8], val.Fulllog[9])

                                }

                        }()




        case "All Probes":

                logtcptracer := make(chan pp.Log, 1)
                        go pp.RunTcptracer(pn[probe], logtcptracer, topResult[0][0])
                        go func() {

                                for val := range logtcptracer {

                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s |IP->%s | SADDR:%s | DADDR:%s | SPORT:%s | DPORT:%s \n",pn[probe],val.fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8],val.Fulllog[9])


                                }

                        }()


                logtcpconnect := make(chan pp.Log, 1)
                        go pp.RunTcpconnect(pn[probe], logtcpconnect, topResult[0][0])
                        go func() {

                                for val := range logtcpconnect {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8])
                                }

                        }()



                logtcpaccept := make(chan pp.Log, 1)
                        go pp.RunTcpaccept(pn[probe], logtcpaccept, topResult[0][0])
                        go func() {

                                for val := range logtcpaccept {

                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | RADDR:%s | RPORT:%s | LADDR:%s | LPORT:%s \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8],val.Fulllog[9])


                                }

                        }()


                logtcplife := make(chan pp.Log, 1)
                        go pp.RunTcplife(pn[probe], logtcplife, topResult[0][0])
                        go func() {

                                for val := range logtcplife {

                                        fmt.Printf("{Probe:%s |Sys_Time: %s |PID:%s | PNAME:%s | LADDRR:%s | LPORT:%s | RADDR:%s | RPORT:%s | TX_KB:%s | RX_KB:%s | MS: %s \n",pn[probe],val.Fulllog[0],val.Fulllog[2],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8],val.Fulllog[9],val.Fulllog[10])

                                }

                        }()

                logexecsnoop := make(chan pp.Log, 1)
                        go pp.RunExecsnoop(pn[probe], logexecsnoop, topResult[0][0])
                        go func() {

                                for val := range logexecsnoop {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s | T:%s | PNAME: %s | PID:%s | PPID:%s | RET:%s | ARGS:%s \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7])

                                }

                        }()

                logbiosnoop := make(chan pp.Log, 1)
                        go pp.RunBiosnoop(pn[probe], logbiosnoop, topResult[0][0])
                        go func() {

                                for val := range logbiosnoop {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s |PNAME: %s | PID:%s | DISK:%s | R/W:%s | SECTOR:%s |BYTES: %s | Lat(ms): %s | \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[2],val.Fulllog[3],val.Fulllog[4],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[9])


                                }

                        }()

                logcachetop := make(chan pp.Log, 1)
                        go pp.RunCachetop(pn[probe], logcachetop, topResult[0][0])
                        go func() {

                                for val := range logcachetop {
                                        fmt.Printf("{Probe:%s |Sys_Time: %s | PID:%s | UID:%s | CMD:%s | HITS:%s | MISS:%s | DIRTIES: %s| READ_HIT%:%s | W_HIT%: %s | \n",pn[probe],val.Fulllog[0],val.Fulllog[1],val.Fulllog[2],val.Fulllog[3],val.Fulllog[5],val.Fulllog[6],val.Fulllog[7],val.Fulllog[8], val.Fulllog[9])

                                }

                        }()


        }







for {

                time.Sleep(time.Duration(1) * time.Second)
}



}

