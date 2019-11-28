# framegram
📷 Add frames to pictures for Instagram from a CLI

## Installation
*framegram* can be installed using golang's *install* command as follows.

```bash
git clone https://github.com/estebanborai/framegram.git
cd ./framegram/

go install
```

## Usage
- [Frame](https://github.com/estebanborai/framegram#frame)
- [Resize](https://github.com/estebanborai/framegram#resize)

### Frame
> Add frames to an Image

Argument | Type | Description
--- | --- | ---
src | `string` | The file to add frames to
out | `string` | The output path

### Resize
> Resizes an Image

Argument | Type | Description
--- | --- | ---
dimensions | `string` | The dimensions to resize the image. Sample: `800x900`
src | `string` | The file to add frames to
out | `string` | The output path
