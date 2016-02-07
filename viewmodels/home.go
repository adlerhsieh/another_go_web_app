package viewmodels

import (

)

type Home struct {
  Title string
  Active string
}

func Get(title, active string) Home {
  result := Home{
    Title: title,
    Active: active,
  }
  return result
}