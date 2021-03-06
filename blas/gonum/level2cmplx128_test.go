// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonum

import (
	"testing"

	"gonum.org/v1/gonum/blas/testblas"
)

func TestZgemv(t *testing.T) {
	testblas.ZgemvTest(t, impl)
}

func TestZgerc(t *testing.T) {
	testblas.ZgercTest(t, impl)
}

func TestZgeru(t *testing.T) {
	testblas.ZgeruTest(t, impl)
}

func TestZhemv(t *testing.T) {
	testblas.ZhemvTest(t, impl)
}

func TestZher(t *testing.T) {
	testblas.ZherTest(t, impl)
}

func TestZher2(t *testing.T) {
	testblas.Zher2Test(t, impl)
}

func TestZhpr(t *testing.T) {
	testblas.ZhprTest(t, impl)
}

func TestZhpr2(t *testing.T) {
	testblas.Zhpr2Test(t, impl)
}

func TestZtrmv(t *testing.T) {
	testblas.ZtrmvTest(t, impl)
}

func TestZtrsv(t *testing.T) {
	testblas.ZtrsvTest(t, impl)
}
