package base

import (
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestRepository_InsertLog(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	test := NewMockLogger(ctr)
	test.EXPECT().InsertData(time.Now(), "opType", []byte{}, true).Return(nil)

}
