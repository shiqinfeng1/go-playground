package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// 写日志文件
	logDir := "./run_log/"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		fmt.Println("Mkdir failed, err:", err)
		return
	}
	fileName := logDir + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	logger := zerolog.New(logFile)

	// 写控制台
	// logger := zerolog.New(os.Stderr)
	log1 := logger.With().Bool("bool", true).Str("string", "with aaa").Logger()
	log1.Info().Msg("log1 ok")
	log1.Debug().Str("string", "debug with bbb").Msg("公共前缀 log1 ok")
	log1.Warn().Str("string", "warn with bbb").Msg("公共前缀log1 ok")
	log1.Error().Str("string", "warn with bbb").Msg("公共前缀log1 ok")
	log1.Error().Str("string", "err with bbb").Err(fmt.Errorf("公共前缀 error"))

	log2 := logger
	log2.Info().Msg("log2 ok")
	log2.Info().Dict("dict", zerolog.Dict().
		Str("bar", "baz").
		Int("n", 1),
	).Msg("嵌套使用")

	log3 := logger.With().Caller().Logger()
	log3.Log().Str("string", "err with ccc").Err(fmt.Errorf("公共前缀 不打印level")).Msg("")
	log3.Info().Str("string", "err with ccc").Err(fmt.Errorf("公共前缀 error")).Send()

	sampled := logger.Sample(&zerolog.BasicSampler{N: 10})

	for i := 0; i < 33; i++ {
		sampled.Info().Str("time", time.Now().Format("2006-01-02 15:04:05")).Msg("采样日志: 每10个日志打印1个")
		time.Sleep(time.Millisecond * 20)
	}

	log.Debug().Err(err).Msg("直接输出")

}
