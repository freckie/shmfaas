// referenced https://medium.com/@juliorenner123/k8s-creating-a-kube-scheduler-plugin-8a826c486a1

package shmscoring

import (
	"context"
	"fmt"

	"github.com/freckie/shmm/shmsched/apis/config"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	klog "k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type ShmScoring struct {
	handle framework.Handle
	conn   *ShmmConnector
}

const Name = "ShmScoring"

var _ = framework.ScorePlugin(&ShmScoring{})

func New(obj runtime.Object, h framework.Handle) (framework.Plugin, error) {
	args, ok := obj.(*config.ShmScoringArgs)
	if !ok {
		return nil, fmt.Errorf("[ShmScoring] want args to be of type ShmScoringArgs, got %T", obj)
	}

	klog.Infof("[ShmScoring] args received. args:%s", args)

	return &ShmScoring{
		handle: h,
		conn:   NewShmmConnector([]string{}, []string{}),
	}, nil
}

func (s *ShmScoring) Name() string {
	return Name
}

func (s *ShmScoring) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status) {
	var score int64

	// getting node metric
	metric, err := s.conn.GetNodeMetric(nodeName)
	if err != nil {
		return 0, framework.NewStatus(framework.Error, fmt.Sprintf("error getting node metrics: %s", err))
	}

	// scoring
	score = metric.ShmSum

	klog.Infof("[ShmScoring] node \"%s\" score \"%d\"", nodeName, score)
	return score, nil
}

func (s *ShmScoring) ScoreExtensions() framework.ScoreExtensions {
	return s
}

func (s *ShmScoring) NormalizeScore(
	ctx context.Context,
	state *framework.CycleState,
	pod *v1.Pod,
	scores framework.NodeScoreList,
) *framework.Status {
	var highestScore int64
	for _, node := range scores {
		if highestScore < node.Score {
			highestScore = node.Score
		}
	}

	for i, node := range scores {
		scores[i].Score = framework.MaxNodeScore - (node.Score * framework.MaxNodeScore / highestScore)
	}

	klog.Infof("[ShmScoring] final scores: %v", scores)
	return nil
}

func main() {
	klog.InfoS("shmscoring")
}
