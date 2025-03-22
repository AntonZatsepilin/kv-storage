box = require('box')

box.cfg{
    listen = 3301,
    log_level = 5,
    log = 'tarantool.log'
}

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

return true