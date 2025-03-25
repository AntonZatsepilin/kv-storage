box.once("bootstrap", function()
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
    
    box.schema.user.create('guest', {if_not_exists = true})
    box.schema.user.grant('guest', 'read,write,execute', 'universe', {if_not_exists = true})
end)