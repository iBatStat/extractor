package http

import (
	_ "github.com/dgrijalva/jwt-go"
)

// we need functions that do the following
//
// 1. generate a new jwt token and stick user claims
//    user claims can be - user id and device type for now
//    for now - expireAt will be 1 hours -- default and hard coded
//    we will also stick the genereated token in the header of the request and redirect to desired site
// 2. authenticate a jwt on a http reqeust
// 3. possibly we want to re-validate a jwt with expire at extended
