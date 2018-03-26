package consolidateBy

import (
	"github.com/go-graphite/carbonapi/expr/helper"
	"github.com/go-graphite/carbonapi/expr/interfaces"
	"github.com/go-graphite/carbonapi/expr/types"
	"github.com/go-graphite/carbonapi/pkg/parser"
)

type consolidateBy struct {
	interfaces.FunctionBase
}

func GetOrder() interfaces.Order {
	return interfaces.Any
}

func New(configFile string) []interfaces.FunctionMetadata {
	res := make([]interfaces.FunctionMetadata, 0)
	f := &consolidateBy{}
	functions := []string{"consolidateBy"}
	for _, n := range functions {
		res = append(res, interfaces.FunctionMetadata{Name: n, F: f})
	}
	return res
}

// consolidateBy(seriesList, aggregationMethod)
func (f *consolidateBy) Do(e parser.Expr, from, until uint32, values map[parser.MetricRequest][]*types.MetricData) ([]*types.MetricData, error) {
	arg, err := helper.GetSeriesArg(e.Args()[0], from, until, values)
	if err != nil {
		return nil, err
	}
	name, err := e.GetStringArg(1)
	if err != nil {
		return nil, err
	}

	var results []*types.MetricData

	for _, a := range arg {
		r := *a

		r.AggregateFunction = types.ConsolidationToFunc[name]

		results = append(results, &r)
	}

	return results, nil
}

// Description is auto-generated description, based on output of https://github.com/graphite-project/graphite-web
func (f *consolidateBy) Description() map[string]types.FunctionDescription {
	return map[string]types.FunctionDescription{
		"consolidateBy": {
			Description: "Takes one metric or a wildcard seriesList and a consolidation function name.\n\nValid function names are 'sum', 'average', 'min', 'max', 'first' & 'last'.\n\nWhen a graph is drawn where width of the graph size in pixels is smaller than\nthe number of datapoints to be graphed, Graphite consolidates the values to\nto prevent line overlap. The consolidateBy() function changes the consolidation\nfunction from the default of 'average' to one of 'sum', 'max', 'min', 'first', or 'last'.\nThis is especially useful in sales graphs, where fractional values make no sense and a 'sum'\nof consolidated values is appropriate.\n\n.. code-block:: none\n\n  &target=consolidateBy(Sales.widgets.largeBlue, 'sum')\n  &target=consolidateBy(Servers.web01.sda1.free_space, 'max')",
			Function:    "consolidateBy(seriesList, consolidationFunc)",
			Group:       "Special",
			Module:      "graphite.render.functions",
			Name:        "consolidateBy",
			Params: []types.FunctionParam{
				{
					Name:     "seriesList",
					Required: true,
					Type:     types.SeriesList,
				},
				{
					Name: "consolidationFunc",
					Options: []string{
						"max",
						"min",
						"sum",
						"average",
						"first",
						"last",
					},
					Required: true,
					Type:     types.String,
				},
			},
		},
	}
}
