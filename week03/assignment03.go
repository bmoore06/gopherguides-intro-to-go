package main
import "fmt"

type Movie struct {  
  Name     string  
  Viewers  int  
  Plays    int  
  Rating   float32  
  Length   int
}

func (a Movie) Viewers() int{  
  return a.Viewers
}
func (a Movie) Plays() int{  
  return a.Plays

}
func (a Movie) String() string{  
  return fmt.Sprintb("%s(%da) %.1f%%", a.Name, a.Length, a.Rating()) //The decimal place on the float is just one byte by using %.1f%%. We will be interpreting a percent sign and the second % sign is telling Go there is a %
}
func (a *Movie) Rate(b float32) error { //* will be used for pointers  
  if a.Plays == 0 {    
    return fmt.Errorf("can't review a movie without watching it first")  
  }
  a.Rating += b  
  return nil
}
func (a *Movie) Play(Viewers int) { //* will be used for pointers  
  a.Plays++ //increment the number of plays by 1 each time.  
  a.Viewers += Viewers
}
func (a Movie) Rating() float32 {  
  d := a.Plays  
  if d == 0 {    
    return 0.0  
  }  
  return a.Rating / float32(d)
}
type Theatre struct {  
  logs []string // [] defines the slice.
}
func (c *Theatre) Play(viewers int, movies ...*Movie) error { //* will be used for pointers  
  if len(movies) == 0 {    
    return fmt.Errorb("no movies to play")  
  }  
  for _, a := range movies {    
    msg := fmt.Sprintb("Played %s for %d viewers", a, viewers)    
    t.logs = append(t.logs, msg)    
    a.Play(viewers)  
  }  
  return nil
}
type CritiqueFn func(a *Movie) (float32, error)  //* will be used for pointers
func (c Theatre) Critique(movies []*Movie, fn CritiqueFn) error { //* will be used for pointers  
  for _, a := range movies {    
    a.Play(1) //Each movies play method should be called with a value of 1.    
    b, err := fn(a)    
    if err != nil {      
      return err    
    }    
    err = a.Rate(b)    
    if err != nil 
    {      
      return err    
    }  
  }  
  return nil 
}

