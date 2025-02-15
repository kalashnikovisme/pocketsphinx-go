// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 05 Sep 2017 19:47:43 MSK.
// By https://git.io/c-for-go. DO NOT EDIT.

package pocketsphinx

/*
#cgo pkg-config: pocketsphinx
#include "pocketsphinx.h"
#include "../sphinxbase/jsgf.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Decoder as declared in pocketsphinx/pocketsphinx.h:67
type Decoder C.ps_decoder_t

// Nbest as declared in pocketsphinx/pocketsphinx.h:74
type Nbest C.ps_nbest_t

// Seg as declared in pocketsphinx/pocketsphinx.h:79
type Seg C.ps_seg_t

// File as declared in include/stdio.h:153
type File C.FILE

// Arg as declared in sphinxbase/cmd_ln.h:93
type Arg struct {
	Name           string
	Type           int32
	Deflt          string
	Doc            string
	ref8495e8ec    *C.arg_t
	allocs8495e8ec interface{}
}

// CommandLn as declared in sphinxbase/cmd_ln.h:179
type CommandLn C.cmd_ln_t

// Anytype as declared in sphinxbase/prim_type.h:112
const sizeofAnytype = unsafe.Sizeof(C.anytype_t{})

type Anytype [sizeofAnytype]byte

// Logmath as declared in sphinxbase/logmath.h:108
type Logmath C.logmath_t

// Fe as declared in sphinxbase/fe.h:287
type Fe C.fe_t

// Feat as declared in sphinxbase/feat.h:162
type Feat C.feat_t

// Lattice as declared in pocketsphinx/ps_lattice.h:62
type Lattice C.ps_lattice_t

// Latnode as declared in pocketsphinx/ps_lattice.h:70
type Latnode C.ps_latnode_t

// LatnodeIter as declared in pocketsphinx/ps_lattice.h:75
type LatnodeIter C.ps_latnode_iter_t

// Latlink as declared in pocketsphinx/ps_lattice.h:83
type Latlink C.ps_latlink_t

// LatlinkIter as declared in pocketsphinx/ps_lattice.h:88
type LatlinkIter C.ps_latlink_iter_t

// NgramModel as declared in sphinxbase/ngram_model.h:66
type NgramModel C.ngram_model_t

// NgramClass as declared in sphinxbase/ngram_model.h:71
type NgramClass C.ngram_class_t

// NgramIter as declared in sphinxbase/ngram_model.h:359
type NgramIter C.ngram_iter_t

// NgramModelSetIter as declared in sphinxbase/ngram_model.h:556
type NgramModelSetIter C.ngram_model_set_iter_t

// Mllr as declared in pocketsphinx/ps_mllr.h:62
type Mllr C.ps_mllr_t

// SearchIter as declared in pocketsphinx/ps_search.h:84
type SearchIter C.ps_search_iter_t

// FsgLink as declared in sphinxbase/fsg_model.h:79
type FsgLink struct {
	FromState      int32
	ToState        int32
	Logs2prob      int32
	Wid            int32
	refe7ca59ff    *C.fsg_link_t
	allocse7ca59ff interface{}
}

// FsgModel as declared in sphinxbase/fsg_model.h:115
type FsgModel C.fsg_model_t

// FsgArciter as declared in sphinxbase/fsg_model.h:130
type FsgArciter C.fsg_arciter_t

// JSGF as declared in sphinxbase/jsgf.h:64
type JSGF C.jsgf_t

// JSGFRule as declared in sphinxbase/jsgf.h:65
type JSGFRule C.jsgf_rule_t

// JSGFRuleIter as declared in sphinxbase/jsgf.h:111
type JSGFRuleIter struct {
	Idx            uint
	refb26723b3    *C.hash_iter_t
	allocsb26723b3 interface{}
}
