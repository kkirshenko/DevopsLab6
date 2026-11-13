package tests

import (
	"context"
	"testing"

	"github.com/kkirshenko/DevopsLab6/internal/restaurant"
)

func TestRestaurantLifecycle(t *testing.T) {
	ctx := context.Background()
	repo := restaurant.NewMemoryRepository()
	svc := restaurant.NewService(repo)
	item := restaurant.Restaurant{ID:"r1",Name:"Cafe",Address:"Main 1",Rating:4.2,Open:true}
	id, err := svc.Create(ctx,item)
	if err != nil || id != "r1" { t.Fatalf("Create: id=%q err=%v",id,err) }
	got, err := svc.GetByID(ctx,id)
	if err != nil || got.Name != item.Name { t.Fatalf("GetByID: %+v %v",got,err) }
	if err := svc.ChangeRating(ctx,id,4.8); err != nil { t.Fatal(err) }
	got, _ = svc.GetByID(ctx,id)
	if got.Rating != 4.8 { t.Fatalf("rating=%v",got.Rating) }
}
func TestValidationAndNotFound(t *testing.T) {
	svc := restaurant.NewService(restaurant.NewMemoryRepository())
	if _,err:=svc.Create(context.Background(),restaurant.Restaurant{ID:"x",Name:"x",Address:"x",Rating:6}); err==nil { t.Fatal("expected validation error") }
	if _,err:=svc.GetByID(context.Background(),"missing"); err!=restaurant.ErrNotFound { t.Fatalf("got %v",err) }
}
