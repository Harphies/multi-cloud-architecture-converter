# Multi-Cloud Architecture Converter: 

To convert architecture diagrams from one cloud provider to another. e.g AWS Reference Architecture diagram to its Azure equvalence etc

[Architecture Converter Components list for Implementation](https://drive.google.com/file/d/1h3RtRcJFahOQA3ERr9A1suhDTpturFXo/view?usp=sharing):
- *decompoer:* is a stateful component that would convert a reference architecture into its component lists and connectors and pass the context to downstream matcher
- *matcher:* get the list of componets from upstream decomper with an extra argument(cloud provider name), search and return the equivalent components for that cloud provider
- *sketcher:* get the component list returned from upstream matcher and the connectors from upstream decomposer then recreate the input reference architecture diagram for the provided cloud provider using the original architecture as a guide for the sketching

Implementation: I'll like this to be implemented in Go as you have capabiliity to pass information around throgh context and much capability for background processess through go routine to fasten the operation and fan out fan in long running processess that would deplay the convertion

## Commands

```
go mod init multi-cloud-architecture-converter
go mod tidy
got get github.com/julienschmidt/httprouter
go get github.com/felixge/httpsnoop
go get github.com/tomasen/realip
go get golang.org/x/time/rate
```