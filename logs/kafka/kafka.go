/*******************************************************
 *  File        :   kafka.go
 *  Author      :   nieaowei
 *  Date        :   2020/1/25 10:57 下午
 *  Notes       :
 *******************************************************/
package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/gogf/gf/frame/g"
	"log-agent-demo/instance"
)

type KafkaInst struct {
	sarama.SyncProducer
	ch chan *instance.Message
}

func init() {

}

func NewKafkaInst(server string, config *sarama.Config) (inst *KafkaInst, err error) {
	inst = &KafkaInst{}
	inst.SyncProducer, err = sarama.NewSyncProducer([]string{server}, config)
	if err != nil {
		return
	}
	return
}

func (p *KafkaInst) Exce() {
	for {
		msg, err := p.ReceMsg()
		if err != nil {
			break
		}
		p.SendMsg(msg)
	}
}

func (p *KafkaInst) LoadConfig() {

}

func (p *KafkaInst) SendMsg(msg *instance.Message) {
	pid, offet, err := p.SendMessage(&sarama.ProducerMessage{
		Topic: msg.Topic,
		Value: sarama.StringEncoder(msg.Text),
	})
	if err != nil {
		return
	}
	g.Log().Info("kafka:", pid, offet)
}

func (p *KafkaInst) ReceMsg() (msg *instance.Message, err error) {
	msg, ok := <-p.ch
	if !ok {
		return nil, errors.New("channel closed.")
	}
	return
}

func DefaultConfig() *sarama.Config {
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	saramaCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	saramaCfg.Producer.Return.Successes = true
	return saramaCfg
}
