package processor;

import(
    "fmt"
    "encoding/json"
	"github.com/wangyihit/thrift_idl/go/thrift_gen/mt/processor"
    "github.com/wangyihit/libmt/util/thrift/serializer"
)

var cmdTemplate = `
console.log('-----');
var data = %s;
var e = %s;
function jsonMap(data, entries){
    try{

for(var j=0; j < entries.length;j++)
        {
            var element = entries[j];
            var src = element.src;
            var dest = element.dest;
            var srcPaths = src.split(".");
            var destPaths = dest.split(".");
            var v = data;
            var i = 0;
            for(i = 0; i < srcPaths.length;i++){
                v = v[srcPaths[i]];
            }
            var k = data;
            for(i = 0; i < destPaths.length - 1; i++){
                if( !data[destPaths[i]] ){
                    data[destPaths[i]] = {};
                }
                k = data[destPaths[i]];
            }
            k[destPaths[destPaths.length - 1]] = v;
        }
        return JSON.stringify(data);
    }catch(e){
        return 1/0;
    }
}
var result = jsonMap(data, e);
console.log("res=" + result);
`;

type JsonFormater struct {
    js Js;
}
var _ Processor = (*Js)(nil)

func NewJsonFormater() *JsonFormater {
    j := &JsonFormater {
        js: NewJS(),
    }
    return j
}


func (c *JsonFormater) Name() string {
	return processor.JsonFormatTask
}
func NewTJsonFormatTask() *processor.TJsonFormatTask {
    task := processor.NewTJsonFormatTask()
    task.FieldMappings = make( []*processor.TJsonFieldMapping)
    return task
}
func (c *JsonFormater) parseTask(data string) (*processor.TJsonFormatTask, error) {
    task := NewTJsonFormatTask()
    err := serializer.BytesToThriftObject((byte[])data, task)
    return task, err
}

func (c *JsonFormater) Run(task *processor.TTask) (*processor.TTaskResult_, error) {
    taskData := task.GetTaskData()
	jsonData := taskData.GetData()
    task, err:= c.parseTask(jsonData)
    jsonStr := task.GetJSONStr()
    entries := task.GetEntries()
    entryString, err := json.Marshal(entries)
    if err != nil {
        return nil, err
    }
    jsCmd := fmt.SPrintf(cmdTemplate, jsonStr, entryString)
	c.vm.Run(jsCmd)
	data, err := c.vm.Get("result")
	if err != nil {
		return nil, err
	}
	taskResult := NewTaskResult_()
	taskResult.Data, err = data.ToString()
	if err != nil {
		taskResult.TaskStatus = processor.TTaskStatus_FAILED
		return taskResult, err
	}
	if taskResult.GetData() == "Infinity" {
		taskResult.TaskStatus = processor.TTaskStatus_FAILED
		return taskResult, errors.New("js code error")
	}
	taskResult.TaskStatus = processor.TTaskStatus_SUCCESS
	return taskResult, nil
}

