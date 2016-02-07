package viewmodels

import (

)

type Categories struct {
  Title string
  Active string
  Categories []Category
}

type Category struct {
  ImageUrl string
  Title string
  Description string
  IsOrientRight bool
}