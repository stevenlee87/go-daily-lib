module example.com/main

go 1.16

require example.com/util v1.0.0
// Warning:(7, 29) Committing local paths might be not portable
replace example.com/util => ../util