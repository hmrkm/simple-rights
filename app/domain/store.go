package domain

//go:generate mockgen -source=$GOFILE -self_package=github.com/hmrkm/simple-rights/$GOPACKAGE -package=$GOPACKAGE -destination=store_mock.go
type Store interface {
	Load(destAddr interface{}, cond interface{}) error
}
