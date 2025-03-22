box = require('box')

local conn = box.cfg.listen ~= nil
if not conn then
    os.exit(1)
end

local ok = pcall(box.space.kv.select, box.space.kv, 'primary', {''}, {limit=1})
os.exit(ok and 0 or 1)