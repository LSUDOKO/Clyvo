package databases

import "errors"

var (
	ErrCantFindProduct        = errors.New("can't find product")
	ErrCantDecodeProducts     = errors.New("can't decode products")
	ErrUserIDNotValid         = errors.New("user id not found")
	ErrorCantUpdateUser       = errors.New("can't update user")
	ErrCantRemoveItemFromCart = errors.New("can't remove item from cart")
	ErrCantGetItem            = errors.New("can't get item from cart")
	ErrCantBuyCartItem        = errors.New("can't buy item from cart")
)

func AddProductToCart() {

}
func RemoveProductFromCart() {

}

func BuyItemFromCart() {

}

func InstantBuyItem() {

}
