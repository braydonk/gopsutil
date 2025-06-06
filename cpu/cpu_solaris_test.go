// SPDX-License-Identifier: BSD-3-Clause
package cpu

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseISAInfo(t *testing.T) {
	cases := []struct {
		filename string
		expected []string
	}{
		{
			"1cpu_1core_isainfo.txt",
			[]string{
				"rdseed", "adx", "avx2", "fma", "bmi2", "bmi1", "rdrand", "f16c", "vmx",
				"avx", "xsave", "pclmulqdq", "aes", "movbe", "sse4.2", "sse4.1", "ssse3", "popcnt",
				"tscp", "cx16", "sse3", "sse2", "sse", "fxsr", "mmx", "cmov", "amd_sysc", "cx8",
				"tsc", "fpu",
			},
		},
		{
			"2cpu_1core_isainfo.txt",
			[]string{
				"rdseed", "adx", "avx2", "fma", "bmi2", "bmi1", "rdrand", "f16c", "vmx",
				"avx", "xsave", "pclmulqdq", "aes", "movbe", "sse4.2", "sse4.1", "ssse3", "popcnt",
				"tscp", "cx16", "sse3", "sse2", "sse", "fxsr", "mmx", "cmov", "amd_sysc", "cx8",
				"tsc", "fpu",
			},
		},
		{
			"2cpu_8core_isainfo.txt",
			[]string{
				"vmx", "avx", "xsave", "pclmulqdq", "aes", "sse4.2", "sse4.1", "ssse3", "popcnt",
				"tscp", "cx16", "sse3", "sse2", "sse", "fxsr", "mmx", "cmov", "amd_sysc", "cx8",
				"tsc", "fpu",
			},
		},
		{
			"2cpu_12core_isainfo.txt",
			[]string{
				"amd_svm", "amd_lzcnt", "popcnt", "amd_sse4a", "tscp", "ahf", "cx16", "sse3", "sse2",
				"sse", "fxsr", "amd_3dnowx", "amd_3dnow", "amd_mmx", "mmx", "cmov", "amd_sysc", "cx8", "tsc", "fpu",
			},
		},
	}

	for _, tc := range cases {
		content, err := os.ReadFile(filepath.Join("testdata", "solaris", tc.filename))
		require.NoErrorf(t, err, "cannot read test case: %s", err)

		sort.Strings(tc.expected)

		flags, err := parseISAInfo(string(content))
		require.NoErrorf(t, err, "parseISAInfo: %s", err)

		require.Truef(t, reflect.DeepEqual(tc.expected, flags), "Bad flags\nExpected: %v\n   Actual: %v", tc.expected, flags)
	}
}

func TestParseProcessorInfo(t *testing.T) {
	cases := []struct {
		filename string
		expected []InfoStat
	}{
		{
			"1cpu_1core_psrinfo.txt",
			[]InfoStat{
				{CPU: 0, VendorID: "GenuineIntel", Family: "6", Model: "78", Stepping: 3, PhysicalID: "0", CoreID: "0", Cores: 1, ModelName: "Intel(r) Core(tm) i7-6567U CPU @ 3.30GHz", Mhz: 3312},
			},
		},
		{
			"2cpu_1core_psrinfo.txt",
			[]InfoStat{
				{CPU: 0, VendorID: "GenuineIntel", Family: "6", Model: "78", Stepping: 3, PhysicalID: "0", CoreID: "0", Cores: 1, ModelName: "Intel(r) Core(tm) i7-6567U CPU @ 3.30GHz", Mhz: 3312},
				{CPU: 1, VendorID: "GenuineIntel", Family: "6", Model: "78", Stepping: 3, PhysicalID: "1", CoreID: "0", Cores: 1, ModelName: "Intel(r) Core(tm) i7-6567U CPU @ 3.30GHz", Mhz: 3312},
			},
		},
		{
			"2cpu_8core_psrinfo.txt",
			[]InfoStat{
				{CPU: 0, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "0", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 1, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "1", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 2, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "2", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 3, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "3", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 4, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "4", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 5, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "5", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 6, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "6", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 7, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "0", CoreID: "7", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 8, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "0", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 9, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "1", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 10, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "2", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 11, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "3", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 12, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "4", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 13, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "5", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 14, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "6", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
				{CPU: 15, VendorID: "GenuineIntel", Family: "6", Model: "45", Stepping: 7, PhysicalID: "1", CoreID: "7", Cores: 2, ModelName: "Intel(r) Xeon(r) CPU E5-2670 0 @ 2.60GHz", Mhz: 2600},
			},
		},
		{
			"2cpu_12core_psrinfo.txt",
			[]InfoStat{
				{CPU: 0, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "0", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 1, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "1", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 2, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "2", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 3, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "3", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 4, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "4", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 5, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "5", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 6, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "6", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 7, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "7", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 8, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "8", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 9, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "9", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 10, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "10", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 11, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "0", CoreID: "11", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 12, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "0", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 13, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "1", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 14, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "2", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 15, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "3", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 16, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "4", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 17, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "5", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 18, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "6", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 19, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "7", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 20, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "8", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 21, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "9", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 22, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "10", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
				{CPU: 23, VendorID: "AuthenticAMD", Family: "16", Model: "9", Stepping: 1, PhysicalID: "1", CoreID: "11", Cores: 1, ModelName: "AMD Opteron(tm) Processor 6176\t[ Socket: G34 ]", Mhz: 2300},
			},
		},
	}

	for _, tc := range cases {
		content, err := os.ReadFile(filepath.Join("testdata", "solaris", tc.filename))
		require.NoErrorf(t, err, "cannot read test case: %s", err)

		cpus, err := parseProcessorInfo(string(content))
		require.NoErrorf(t, err, "cannot parse processor info: %s", err)

		require.Truef(t, reflect.DeepEqual(tc.expected, cpus), "Bad Processor Info\nExpected: %v\n   Actual: %v", tc.expected, cpus)
	}
}
