# Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: circint01

debug:
	go test -gcflags "-N -l" -c -o /tmp/test_eslopes_geo

circint01: debug
	gdb --args /tmp/test_eslopes_geo -test.run=circint01
