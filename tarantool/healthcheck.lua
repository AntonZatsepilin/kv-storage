local function check_kv_space()
    if not box.space.kv then
        return false, "KV space not exists"
    end
    return true
end

local conn_ok, conn_err = pcall(box.cfg, {})
local space_ok, space_err = check_kv_space()

if not conn_ok or not space_ok then
    os.exit(1)
end
os.exit(0)