package assignment03

import "fmt"

type Movie struct {
  
  
}

func (m Movie) Viewers() int {
  return m.viewers
  
}

func (m Movie) Plays() int {
  return m.plays
}

func (m Movie) Rating() float32 {
  a := m.plays
  if a = 0 {
    return 0.0
  }
  return m.rating / float32(a)
}

func (m Movie) String() string {
  return fmt.Sprintf(%s (%dm) %.1f%%, m.
  
}

func (m *Movie) Rate(f float32) error {
  if m.plays = 0 {
    return fmt.Errorf("cant review
}

func (m Movie) Play(viewers int) {
  
}

type Theatre struct {
  
}

func (t *Theatre) Play(viewers int, movies ... *Movie) error {
  
}

type CritiqueFn func(m * Movie) (float32, error)

func (t Theatre) Critique(movies []*Movie, fn CritiqueFn) error {
  
}


