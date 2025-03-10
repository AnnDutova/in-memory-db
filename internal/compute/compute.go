package compute

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"go.uber.org/zap"

	"github.com/AnnDutova/in-memory-db/internal/engine"
)

type (
	computeImpl struct {
		logger *zap.Logger
		engine engine.Engine
	}
)

func New(logger *zap.Logger) (Compute, error) {
	eng, err := engine.New(logger)
	if err != nil {
		return nil, err
	}

	return &computeImpl{
		engine: eng,
		logger: logger,
	}, nil
}

func (p *computeImpl) Parse(_ context.Context, in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		cmd, args, err := p.parseLine(strings.Fields(line))
		if err != nil {
			return fmt.Errorf("failed to parse line '%s': %w", line, err)
		}

		switch cmd {
		case GetCommand:
			value, err := p.handleGet(args)
			if err != nil {
				p.logger.Error("[error]", zap.Error(err))
				return err
			}
			p.logger.Info("[ok]", zap.String(args[0], value))

			if _, err := out.Write([]byte(value)); err != nil {
				return err
			}
			p.logger.Info("return value to stdout", zap.String(args[0], value))
		case SetCommand:
			if err := p.handleSet(args); err != nil {
				p.logger.Error("[error]", zap.Error(err))
				return err
			}
		case DelCommand:
			if err := p.handleDel(args); err != nil {
				p.logger.Error("[error]", zap.Error(err))
				return err
			}
		}
	}
	return nil
}

func (p *computeImpl) parseLine(in []string) (Command, Arguments, error) {
	if len(in) == 0 {
		return "", nil, ErrEmptyQuerry
	}

	cmd := in[0]
	if err := new(Command).validate(cmd); err != nil {
		return "", nil, err
	}

	args := in[1:]
	if err := new(Arguments).validate(Command(cmd), args); err != nil {
		return "", nil, err
	}

	return Command(cmd), Arguments(args), nil
}

func (p *computeImpl) handleGet(args []string) (string, error) {
	value, err := p.engine.Get(args[0])
	if err != nil {
		return "", fmt.Errorf("failed to execute command GET: %w", err)
	}
	return value, nil
}

func (p *computeImpl) handleSet(args []string) error {
	if err := p.engine.Set(args[0], args[1]); err != nil {
		return fmt.Errorf("failed to execute command SET: %w", err)
	}
	p.logger.Info("[ok] set", zap.String(args[0], args[1]))
	return nil
}

func (p *computeImpl) handleDel(args []string) error {
	if err := p.engine.Delite(args[0]); err != nil {
		return fmt.Errorf("failed to execute command DEL: %w", err)
	}
	p.logger.Info("[ok] del", zap.String(args[0], ""))
	return nil
}
