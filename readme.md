# Go GraphQL

[query={artists{id,name,type}}](http://localhost:8088/graphql?query={artists{id,name,type}})

[query={albums(id:"lz-led-zeppelin"){id,artist,title,year,type}}](http://localhost:8088/graphql?query={albums(id:"lz-led-zeppelin"){id,artist,title,year,type}})

[query={songs(album:"lz-led-zeppelin"){id,album,title,duration,type}}](http://localhost:8088/graphql?query={songs(album:"lz-led-zeppelin"){id,album,title,duration,type}})
