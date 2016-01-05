package parser

import "github.com/xlab/c/cc"

type TargetArchBits int

const (
	Arch32 TargetArchBits = 32
	Arch64 TargetArchBits = 64
)

var predefinedBase = `
#define __STDC_HOSTED__ 1
#define __STDC_VERSION__ 199901L
#define __STDC__ 1
#define __GNUC__ 4
#define __builtin_va_list void *
#define __asm__(x)
#define __asm(x)
#define __inline inline
#define __signed signed
#define __attribute__(x)
#define __POSIX_C_DEPRECATED(ver)
`

var predefines = map[TargetArchBits]string{
	Arch32: predefinedBase + `#define __i386__ 1`,
	Arch64: predefinedBase + `#define __x86_64__ 1`,
}

var models = map[TargetArchBits]*cc.Model{
	Arch32: model32,
	Arch64: model64,
}

var arches = map[string]TargetArchBits{
	"386":    Arch32,
	"arm":    Arch32,
	"armbe":  Arch32,
	"mips":   Arch32,
	"mipsle": Arch32,
	"sparc":  Arch32,
	//
	"amd64":       Arch64,
	"amd64p32":    Arch64,
	"arm64":       Arch64,
	"arm64be":     Arch64,
	"ppc64":       Arch64,
	"ppc64le":     Arch64,
	"mips64":      Arch64,
	"mips64le":    Arch64,
	"mips64p32":   Arch64,
	"mips64p32le": Arch64,
	"sparc64":     Arch64,
}

var model32 = &cc.Model{
	Items: map[cc.Kind]cc.ModelItem{
		cc.Ptr:               {4, 4, "__TODO_PTR"},
		cc.Void:              {0, 1, "__TODO_VOID"},
		cc.Char:              {1, 1, "int8"},
		cc.SChar:             {1, 1, "int8"},
		cc.UChar:             {1, 1, "byte"},
		cc.Short:             {2, 2, "int16"},
		cc.UShort:            {2, 2, "uint16"},
		cc.Int:               {4, 4, "int32"},
		cc.UInt:              {4, 4, "uint32"},
		cc.Long:              {4, 4, "int32"},
		cc.ULong:             {4, 4, "uint32"},
		cc.LongLong:          {8, 8, "int64"},
		cc.ULongLong:         {8, 8, "uint64"},
		cc.Float:             {4, 4, "float32"},
		cc.Double:            {8, 8, "float64"},
		cc.LongDouble:        {8, 8, "float64"},
		cc.Bool:              {1, 1, "bool"},
		cc.FloatComplex:      {8, 8, "complex64"},
		cc.DoubleComplex:     {16, 16, "complex128"},
		cc.LongDoubleComplex: {16, 16, "complex128"},
	},
}

var model64 = &cc.Model{
	Items: map[cc.Kind]cc.ModelItem{
		cc.Ptr:               {8, 8, "__TODO_PTR"},
		cc.Void:              {0, 1, "__TODO_VOID"},
		cc.Char:              {1, 1, "int8"},
		cc.SChar:             {1, 1, "int8"},
		cc.UChar:             {1, 1, "byte"},
		cc.Short:             {2, 2, "int16"},
		cc.UShort:            {2, 2, "uint16"},
		cc.Int:               {4, 4, "int32"},
		cc.UInt:              {4, 4, "uint32"},
		cc.Long:              {8, 8, "int64"},
		cc.ULong:             {8, 8, "uint64"},
		cc.LongLong:          {8, 8, "int64"},
		cc.ULongLong:         {8, 8, "uint64"},
		cc.Float:             {4, 4, "float32"},
		cc.Double:            {8, 8, "float64"},
		cc.LongDouble:        {8, 8, "float64"},
		cc.Bool:              {1, 1, "bool"},
		cc.FloatComplex:      {8, 8, "complex64"},
		cc.DoubleComplex:     {16, 16, "complex128"},
		cc.LongDoubleComplex: {16, 16, "complex128"},
	},
}
