// automatically generated by the FlatBuffers compiler, do not modify

import * as flatbuffers from 'flatbuffers';

export class PredictScoreResponse {
  bb: flatbuffers.ByteBuffer|null = null;
  bb_pos = 0;
  __init(i:number, bb:flatbuffers.ByteBuffer):PredictScoreResponse {
  this.bb_pos = i;
  this.bb = bb;
  return this;
}

static getRootAsPredictScoreResponse(bb:flatbuffers.ByteBuffer, obj?:PredictScoreResponse):PredictScoreResponse {
  return (obj || new PredictScoreResponse()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

static getSizePrefixedRootAsPredictScoreResponse(bb:flatbuffers.ByteBuffer, obj?:PredictScoreResponse):PredictScoreResponse {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new PredictScoreResponse()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

afc():number {
  const offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? this.bb!.readInt32(this.bb_pos + offset) : 0;
}

mutate_afc(value:number):boolean {
  const offset = this.bb!.__offset(this.bb_pos, 4);

  if (offset === 0) {
    return false;
  }

  this.bb!.writeInt32(this.bb_pos + offset, value);
  return true;
}

nfc():number {
  const offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.readInt32(this.bb_pos + offset) : 0;
}

mutate_nfc(value:number):boolean {
  const offset = this.bb!.__offset(this.bb_pos, 6);

  if (offset === 0) {
    return false;
  }

  this.bb!.writeInt32(this.bb_pos + offset, value);
  return true;
}

error():string|null
error(optionalEncoding:flatbuffers.Encoding):string|Uint8Array|null
error(optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 8);
  return offset ? this.bb!.__string(this.bb_pos + offset, optionalEncoding) : null;
}

static startPredictScoreResponse(builder:flatbuffers.Builder) {
  builder.startObject(3);
}

static addAfc(builder:flatbuffers.Builder, afc:number) {
  builder.addFieldInt32(0, afc, 0);
}

static addNfc(builder:flatbuffers.Builder, nfc:number) {
  builder.addFieldInt32(1, nfc, 0);
}

static addError(builder:flatbuffers.Builder, errorOffset:flatbuffers.Offset) {
  builder.addFieldOffset(2, errorOffset, 0);
}

static endPredictScoreResponse(builder:flatbuffers.Builder):flatbuffers.Offset {
  const offset = builder.endObject();
  return offset;
}

static createPredictScoreResponse(builder:flatbuffers.Builder, afc:number, nfc:number, errorOffset:flatbuffers.Offset):flatbuffers.Offset {
  PredictScoreResponse.startPredictScoreResponse(builder);
  PredictScoreResponse.addAfc(builder, afc);
  PredictScoreResponse.addNfc(builder, nfc);
  PredictScoreResponse.addError(builder, errorOffset);
  return PredictScoreResponse.endPredictScoreResponse(builder);
}
}
