// Code generated by tinyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	state "github.com/CosmWasm/cosmwasm-go/example/voter/src/state"
	types "github.com/CosmWasm/cosmwasm-go/std/types"
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes(in *jlexer.Lexer, out *VoteTally) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "option":
			out.Option = string(in.String())
		case "total_yes":
			out.TotalYes = uint32(in.Uint32())
		case "total_no":
			out.TotalNo = uint32(in.Uint32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes(out *jwriter.Writer, in VoteTally) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"option\":"
		out.RawString(prefix[1:])
		out.String(string(in.Option))
	}
	{
		const prefix string = ",\"total_yes\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.TotalYes))
	}
	{
		const prefix string = ",\"total_no\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.TotalNo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VoteTally) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v VoteTally) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VoteTally) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *VoteTally) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes1(in *jlexer.Lexer, out *QueryVotingResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		case "name":
			out.Name = string(in.String())
		case "creator_addr":
			out.CreatorAddr = string(in.String())
		case "start_time":
			out.StartTime = uint64(in.Uint64())
		case "end_time":
			out.EndTime = uint64(in.Uint64())
		case "tallies":
			if in.IsNull() {
				in.Skip()
				out.Tallies = nil
			} else {
				in.Delim('[')
				if out.Tallies == nil {
					if !in.IsDelim(']') {
						out.Tallies = make([]state.Tally, 0, 1)
					} else {
						out.Tallies = []state.Tally{}
					}
				} else {
					out.Tallies = (out.Tallies)[:0]
				}
				for !in.IsDelim(']') {
					var v1 state.Tally
					(v1).UnmarshalTinyJSON(in)
					out.Tallies = append(out.Tallies, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes1(out *jwriter.Writer, in QueryVotingResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"creator_addr\":"
		out.RawString(prefix)
		out.String(string(in.CreatorAddr))
	}
	{
		const prefix string = ",\"start_time\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.StartTime))
	}
	{
		const prefix string = ",\"end_time\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.EndTime))
	}
	{
		const prefix string = ",\"tallies\":"
		out.RawString(prefix)
		if in.Tallies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tallies {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalTinyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryVotingResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v QueryVotingResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryVotingResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes1(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *QueryVotingResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes1(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes2(in *jlexer.Lexer, out *QueryVotingRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes2(out *jwriter.Writer, in QueryVotingRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryVotingRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v QueryVotingRequest) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryVotingRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes2(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *QueryVotingRequest) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes2(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes3(in *jlexer.Lexer, out *QueryTallyResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "open":
			out.Open = bool(in.Bool())
		case "votes":
			if in.IsNull() {
				in.Skip()
				out.Votes = nil
			} else {
				in.Delim('[')
				if out.Votes == nil {
					if !in.IsDelim(']') {
						out.Votes = make([]VoteTally, 0, 2)
					} else {
						out.Votes = []VoteTally{}
					}
				} else {
					out.Votes = (out.Votes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 VoteTally
					(v4).UnmarshalTinyJSON(in)
					out.Votes = append(out.Votes, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes3(out *jwriter.Writer, in QueryTallyResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"open\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Open))
	}
	{
		const prefix string = ",\"votes\":"
		out.RawString(prefix)
		if in.Votes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Votes {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalTinyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryTallyResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v QueryTallyResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryTallyResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes3(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *QueryTallyResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes3(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes4(in *jlexer.Lexer, out *QueryTallyRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes4(out *jwriter.Writer, in QueryTallyRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryTallyRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v QueryTallyRequest) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryTallyRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes4(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *QueryTallyRequest) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes4(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes5(in *jlexer.Lexer, out *QueryReleaseStatsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "count":
			out.Count = uint64(in.Uint64())
		case "total_amount":
			if in.IsNull() {
				in.Skip()
				out.TotalAmount = nil
			} else {
				in.Delim('[')
				if out.TotalAmount == nil {
					if !in.IsDelim(']') {
						out.TotalAmount = make([]types.Coin, 0, 2)
					} else {
						out.TotalAmount = []types.Coin{}
					}
				} else {
					out.TotalAmount = (out.TotalAmount)[:0]
				}
				for !in.IsDelim(']') {
					var v7 types.Coin
					(v7).UnmarshalTinyJSON(in)
					out.TotalAmount = append(out.TotalAmount, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes5(out *jwriter.Writer, in QueryReleaseStatsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Count))
	}
	{
		const prefix string = ",\"total_amount\":"
		out.RawString(prefix)
		if in.TotalAmount == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.TotalAmount {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalTinyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryReleaseStatsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v QueryReleaseStatsResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryReleaseStatsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes5(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *QueryReleaseStatsResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes5(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes6(in *jlexer.Lexer, out *QueryParamsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "owner_addr":
			out.OwnerAddr = string(in.String())
		case "new_voting_cost":
			(out.NewVotingCost).UnmarshalTinyJSON(in)
		case "vote_cost":
			(out.VoteCost).UnmarshalTinyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes6(out *jwriter.Writer, in QueryParamsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"owner_addr\":"
		out.RawString(prefix[1:])
		out.String(string(in.OwnerAddr))
	}
	{
		const prefix string = ",\"new_voting_cost\":"
		out.RawString(prefix)
		(in.NewVotingCost).MarshalTinyJSON(out)
	}
	{
		const prefix string = ",\"vote_cost\":"
		out.RawString(prefix)
		(in.VoteCost).MarshalTinyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryParamsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v QueryParamsResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryParamsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes6(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *QueryParamsResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes6(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes7(in *jlexer.Lexer, out *QueryOpenResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ids":
			if in.IsNull() {
				in.Skip()
				out.Ids = nil
			} else {
				in.Delim('[')
				if out.Ids == nil {
					if !in.IsDelim(']') {
						out.Ids = make([]uint64, 0, 8)
					} else {
						out.Ids = []uint64{}
					}
				} else {
					out.Ids = (out.Ids)[:0]
				}
				for !in.IsDelim(']') {
					var v10 uint64
					v10 = uint64(in.Uint64())
					out.Ids = append(out.Ids, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes7(out *jwriter.Writer, in QueryOpenResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ids\":"
		out.RawString(prefix[1:])
		if in.Ids == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Ids {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueryOpenResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v QueryOpenResponse) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueryOpenResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes7(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *QueryOpenResponse) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes7(l, v)
}
func tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes8(in *jlexer.Lexer, out *MsgQuery) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "params":
			if in.IsNull() {
				in.Skip()
				out.Params = nil
			} else {
				if out.Params == nil {
					out.Params = new(struct{})
				}
				tinyjson8354aa0cDecode(in, out.Params)
			}
		case "voting":
			if in.IsNull() {
				in.Skip()
				out.Voting = nil
			} else {
				if out.Voting == nil {
					out.Voting = new(QueryVotingRequest)
				}
				(*out.Voting).UnmarshalTinyJSON(in)
			}
		case "tally":
			if in.IsNull() {
				in.Skip()
				out.Tally = nil
			} else {
				if out.Tally == nil {
					out.Tally = new(QueryTallyRequest)
				}
				(*out.Tally).UnmarshalTinyJSON(in)
			}
		case "open":
			if in.IsNull() {
				in.Skip()
				out.Open = nil
			} else {
				if out.Open == nil {
					out.Open = new(struct{})
				}
				tinyjson8354aa0cDecode(in, out.Open)
			}
		case "release_stats":
			if in.IsNull() {
				in.Skip()
				out.ReleaseStats = nil
			} else {
				if out.ReleaseStats == nil {
					out.ReleaseStats = new(struct{})
				}
				tinyjson8354aa0cDecode(in, out.ReleaseStats)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes8(out *jwriter.Writer, in MsgQuery) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Params != nil {
		const prefix string = ",\"params\":"
		first = false
		out.RawString(prefix[1:])
		tinyjson8354aa0cEncode(out, *in.Params)
	}
	if in.Voting != nil {
		const prefix string = ",\"voting\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Voting).MarshalTinyJSON(out)
	}
	if in.Tally != nil {
		const prefix string = ",\"tally\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Tally).MarshalTinyJSON(out)
	}
	if in.Open != nil {
		const prefix string = ",\"open\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		tinyjson8354aa0cEncode(out, *in.Open)
	}
	if in.ReleaseStats != nil {
		const prefix string = ",\"release_stats\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		tinyjson8354aa0cEncode(out, *in.ReleaseStats)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MsgQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v MsgQuery) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson8354aa0cEncodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MsgQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes8(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *MsgQuery) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson8354aa0cDecodeGithubComCosmWasmCosmwasmGoExampleVoterSrcTypes8(l, v)
}
func tinyjson8354aa0cDecode(in *jlexer.Lexer, out *struct{}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func tinyjson8354aa0cEncode(out *jwriter.Writer, in struct{}) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
