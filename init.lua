box.cfg{
    listen = '0.0.0.0:3301',
    wal_mode = 'none',
    memtx_memory = 256 * 1024 * 1024
}
box.schema.space.create('kv', {if_not_exists = true})
box.space.kv:format({
    {name = 'key', type = 'string'},
    {name = 'value', type = '*'}
})
box.space.kv:create_index('primary', {parts = {'key'}, if_not_exists = true})
box.schema.user.create('Anton', {password = '12345', if_not_exists = true})
box.schema.user.grant('admin', 'read,write,execute', 'universe', nil, {if_not_exists = true})


-- Create a space --
box.schema.space.create('bands')

-- Specify field names and types --
box.space.bands:format({
    { name = 'id', type = 'unsigned' },
    { name = 'band_name', type = 'string' },
    { name = 'year', type = 'unsigned' }
})

-- Create indexes --
box.space.bands:create_index('primary', { parts = { 'id' } })
box.space.bands:create_index('band', { parts = { 'band_name' } })
box.space.bands:create_index('year_band', { parts = { { 'year' }, { 'band_name' } } })

-- Create a stored function --
box.schema.func.create('get_bands_older_than', {
    body = [[
    function(year)
        return box.space.bands.index.year_band:select({ year }, { iterator = 'LT', limit = 10 })
    end
    ]]
})