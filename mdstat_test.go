// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package procfs

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFS_MDStat(t *testing.T) {
	fs := getProcFixtures(t)
	mdStats, err := fs.MDStat()

	if err != nil {
		t.Fatalf("parsing of reference-file failed entirely: %s", err)
	}

	refs := map[string]MDStat{
		"md127": {Name: "md127", ActivityState: "active", DisksActive: 2, DisksTotal: 2, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 312319552, BlocksSynced: 312319552, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdi2", "sdj2"}},
		"md0":   {Name: "md0", ActivityState: "active", DisksActive: 2, DisksTotal: 2, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 248896, BlocksSynced: 248896, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdi1", "sdj1"}},
		"md4":   {Name: "md4", ActivityState: "inactive", DisksActive: 0, DisksTotal: 0, DisksFailed: 1, DisksDown: 0, DisksSpare: 1, BlocksTotal: 4883648, BlocksSynced: 4883648, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sda3", "sdb3"}},
		"md6":   {Name: "md6", ActivityState: "recovering", DisksActive: 1, DisksTotal: 2, DisksFailed: 1, DisksDown: 1, DisksSpare: 1, BlocksTotal: 195310144, BlocksSynced: 16775552, BlocksSyncedPct: 8.5, BlocksSyncedFinishTime: 17, BlocksSyncedSpeed: 259783, Devices: []string{"sdb2", "sdc", "sda2"}},
		"md3":   {Name: "md3", ActivityState: "active", DisksActive: 8, DisksTotal: 8, DisksFailed: 0, DisksDown: 0, DisksSpare: 2, BlocksTotal: 5853468288, BlocksSynced: 5853468288, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sda1", "sdh1", "sdg1", "sdf1", "sde1", "sdd1", "sdc1", "sdb1", "sdd1", "sdd2"}},
		"md8":   {Name: "md8", ActivityState: "resyncing", DisksActive: 2, DisksTotal: 2, DisksFailed: 0, DisksDown: 0, DisksSpare: 2, BlocksTotal: 195310144, BlocksSynced: 16775552, BlocksSyncedPct: 8.5, BlocksSyncedFinishTime: 17, BlocksSyncedSpeed: 259783, Devices: []string{"sdb1", "sda1", "sdc", "sde"}},
		"md7":   {Name: "md7", ActivityState: "active", DisksActive: 3, DisksTotal: 4, DisksFailed: 1, DisksDown: 1, DisksSpare: 0, BlocksTotal: 7813735424, BlocksSynced: 7813735424, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdb1", "sde1", "sdd1", "sdc1"}},
		"md9":   {Name: "md9", ActivityState: "resyncing", DisksActive: 4, DisksTotal: 4, DisksSpare: 1, DisksDown: 0, DisksFailed: 2, BlocksTotal: 523968, BlocksSynced: 0, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdc2", "sdd2", "sdb2", "sda2", "sde", "sdf", "sdg"}},
		"md10":  {Name: "md10", ActivityState: "active", DisksActive: 2, DisksTotal: 2, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 314159265, BlocksSynced: 314159265, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sda1", "sdb1"}},
		"md11":  {Name: "md11", ActivityState: "resyncing", DisksActive: 2, DisksTotal: 2, DisksFailed: 1, DisksDown: 0, DisksSpare: 2, BlocksTotal: 4190208, BlocksSynced: 0, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdb2", "sdc2", "sdc3", "hda", "ssdc2"}},
		"md12":  {Name: "md12", ActivityState: "active", DisksActive: 2, DisksTotal: 2, DisksSpare: 0, DisksDown: 0, DisksFailed: 0, BlocksTotal: 3886394368, BlocksSynced: 3886394368, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdc2", "sdd2"}},
		"md120": {Name: "md120", ActivityState: "active", DisksActive: 2, DisksTotal: 2, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 2095104, BlocksSynced: 2095104, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sda1", "sdb1"}},
		"md126": {Name: "md126", ActivityState: "active", DisksActive: 2, DisksTotal: 2, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 1855870976, BlocksSynced: 1855870976, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdb", "sdc"}},
		"md219": {Name: "md219", ActivityState: "inactive", DisksTotal: 0, DisksFailed: 0, DisksActive: 0, DisksDown: 0, DisksSpare: 3, BlocksTotal: 7932, BlocksSynced: 7932, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdc", "sda"}},
		"md00":  {Name: "md00", ActivityState: "active", DisksActive: 1, DisksTotal: 1, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 4186624, BlocksSynced: 4186624, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"xvdb"}},
		"md101": {Name: "md101", ActivityState: "active", DisksActive: 3, DisksTotal: 3, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 322560, BlocksSynced: 322560, BlocksSyncedPct: 0, BlocksSyncedFinishTime: 0, BlocksSyncedSpeed: 0, Devices: []string{"sdb", "sdd", "sdc"}},
		"md201": {Name: "md201", ActivityState: "checking", DisksActive: 2, DisksTotal: 2, DisksFailed: 0, DisksDown: 0, DisksSpare: 0, BlocksTotal: 1993728, BlocksSynced: 114176, BlocksSyncedPct: 5.7, BlocksSyncedFinishTime: 0.2, BlocksSyncedSpeed: 114176, Devices: []string{"sda3", "sdb3"}},
	}

	if want, have := len(refs), len(mdStats); want != have {
		t.Errorf("want %d parsed md-devices, have %d", want, have)
	}
	for _, md := range mdStats {
		if want, have := refs[md.Name], md; !cmp.Equal(want, have) {
			t.Errorf("%s: want %v, have %v", md.Name, want, have)
		}
	}

}

func TestMdstatScaledProgress(t *testing.T) {

	hugeArrayMdstat := []byte(`
Personalities : [raid6] [raid5] [raid4]
md0 : active raid6 sdj1[9] sdi1[8] sdh1[7] sde1[4] sdg1[6] sda1[0] sdb1[1] sdc1[2] sdf1[5] sdd1[3]
      140617179136 blocks super 1.2 level 6, 512k chunk, algorithm 2 [10/10] [UUUUUUUUUU]
      [==========>..........]  resync = 51.6% (9083530032/17577147392) finish=732.9min speed=193145K/sec
      bitmap: 64/131 pages [256KB], 65536KB chunk

unused devices: <none>
`)

	mdStats, err := parseMDStat(hugeArrayMdstat)

	if err != nil {
		t.Fatalf("parsing of mdstat failed entirely: %s", err)
	}

	progressScale := int64(140617179136) / int64(17577147392)
	refs := map[string]MDStat{
		"md0": {
			Name:                   "md0",
			ActivityState:          "resyncing",
			DisksActive:            10,
			DisksTotal:             10,
			DisksFailed:            0,
			DisksDown:              0,
			DisksSpare:             0,
			BlocksTotal:            140617179136,
			BlocksSynced:           9083530032 * progressScale,
			BlocksSyncedPct:        51.6,
			BlocksSyncedFinishTime: 732.9,
			BlocksSyncedSpeed:      193145,
			Devices:                []string{"sdj1", "sdi1", "sdh1", "sde1", "sdg1", "sda1", "sdb1", "sdc1", "sdf1", "sdd1"},
		},
	}

	if want, have := len(refs), len(mdStats); want != have {
		t.Errorf("want %d parsed md-devices, have %d", want, have)
	}
	for _, md := range mdStats {
		if want, have := refs[md.Name], md; !cmp.Equal(want, have) {
			t.Errorf("%s: want %v, have %v", md.Name, want, have)
		}
		if havePct := 100.0 * float64(md.BlocksSynced) / float64(md.BlocksTotal); math.Abs(havePct-md.BlocksSyncedPct) > 0.1 {
			t.Errorf("%s: synced pct %v, have %v", md.Name, md.BlocksSyncedPct, havePct)
		}
	}

}

func TestInvalidMdstat(t *testing.T) {
	invalidMount := [][]byte{[]byte(`
Personalities : [invalid]
md3 : invalid
      314159265 blocks 64k chunks

unused devices: <none>
`),
		[]byte(`
md12 : active raid0 sdc2[0] sdd2[1]

      3886394368 blocks super 1.2 512k chunks
`)}

	for _, invalid := range invalidMount {
		_, err := parseMDStat(invalid)
		if err == nil {
			t.Fatalf("parsing of invalid reference file did not find any errors")
		}
	}
}
