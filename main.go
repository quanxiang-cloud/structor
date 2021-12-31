package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/quanxiang-cloud/structor/internal/dorm"
	"github.com/quanxiang-cloud/structor/internal/dorm/clause"
)

func main() {
	flag.Parse()

	d, err := dorm.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()

	// expr, err := clause.New().GetExpression("sum", "types", "types")

	id, err := clause.New().GetExpression("like", "_id", "00")
	if err != nil {
		fmt.Println(err)
		return
	}

	app, err := clause.New().GetExpression("eq", "app_id", "q88h4")
	if err != nil {
		fmt.Println(err)
		return
	}

	expr, err := clause.New().GetExpression("and", "", id, app)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := d.Table("permission_group").Find(ctx, expr, clause.FindOptions{
		Page: 1,
		Size: 2,
	})
	fmt.Println(data, err)
}
