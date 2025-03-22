box = require('box')
http = require('http.server')

box.cfg{
    listen = 3301,
    log_level = 5,
    log = 'tarantool.log'
}

-- Create space and indexes (existing code)
box.schema.space.create('kv', {
    if_not_exists = true,
    format = {
        {name = 'key', type = 'string'},
        {name = 'value', type = '*'}
    }
})

box.space.kv:create_index('primary', {
    type = 'hash',
    parts = {'key'},
    if_not_exists = true
})

-- Start HTTP server
http.new('0.0.0.0', 8081):start()

-- Grant permissions (existing code)
box.schema.user.grant('guest', 'read,write,execute', 'space', 'kv')
box.schema.user.grant('guest', 'create,drop', 'space')
box.schema.user.grant('guest', 'read,write', 'universe')

return true