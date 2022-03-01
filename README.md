# learn-go-lang

# Git Commands
git remote add origin https://github.com/emuhlestein/learn-go-lang.git  
git branch -m main  
git branch --set-upstream-to=origin/main main  
git pull --allow-unrelated-histories  

GO is statically typed. 
All variables are initilized.
Zero Values:
- numeric types: 0
- bool type: false
- string type: "" (empty string)
- pointer type: nill

// you must provide a type for each variable you declare or Go should be able to infer it

age := 18
_ = age // blank variable removes unused var error

