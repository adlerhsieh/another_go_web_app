package viewmodels

import (

)

type Home struct {
  Title string
  Active string
}

func GetHome() Home {
  result := Home{
    Title: "Lemon",
    Active: "index",
  }
  return result
}