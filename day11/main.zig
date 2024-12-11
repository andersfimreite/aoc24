const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buffer: [64]u8 = undefined;
    const bytes_read = try file.read(buffer[0..]);

    const res1 = try part1(&buffer[0..bytes_read]);
    std.debug.print("Part 1: {}\n", .{res1});

    const res2 = try part2(&buffer[0..bytes_read]);
    std.debug.print("Part 2: {}\n", .{res2});
}

pub fn part1(data: *const []u8) !usize {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    var stones = std.ArrayList(u64).init(allocator);
    defer stones.deinit();

    var stones2 = std.ArrayList(u64).init(allocator);
    defer stones2.deinit();

    const trimmed = std.mem.trim(u8, data.*, "\n");
    var it = std.mem.split(u8, trimmed, " ");
    while (it.next()) |part| {
        const num = try std.fmt.parseInt(u64, part, 10);
        try stones.append(num);
    }

    const blinks: u64 = 25;
    for (0..blinks) |_| {
        for (stones.items) |num| {
            if (num == 0) {
                try stones2.append(1);
                continue;
            }

            const digits = @as(u64, @intFromFloat(@log10(@as(f64, @floatFromInt(num))) + 1));
            if (@mod(digits, 2) == 0) {
                const tens = std.math.pow(u64, 10, digits / 2);
                try stones2.append(num / tens);
                try stones2.append(@mod(num, tens));
                continue;
            }

            try stones2.append(num * 2024);
        }

        const tmp = stones;
        stones = stones2;
        stones2 = tmp;
        stones2.clearRetainingCapacity();
    }

    return stones.items.len;
}

pub fn part2(data: *const []u8) !u64 {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    var stones = std.AutoHashMap(u64, u64).init(allocator);
    defer stones.deinit();

    var stones2 = std.AutoHashMap(u64, u64).init(allocator);
    defer stones2.deinit();

    const trimmed = std.mem.trim(u8, data.*, "\n");
    var it = std.mem.split(u8, trimmed, " ");
    while (it.next()) |part| {
        const num = try std.fmt.parseInt(u64, part, 10);
        try stones.put(num, 1);
    }

    const blinks: u64 = 75;
    for (0..blinks) |_| {
        var itr = stones.iterator();
        while (itr.next()) |kv| {
            const num = kv.key_ptr.*;
            const cnt = kv.value_ptr.*;

            if (num == 0) {
                if (stones2.get(1)) |c| {
                    try stones2.put(1, c + cnt);
                } else {
                    try stones2.put(1, cnt);
                }
                continue;
            }

            const digits = @as(u64, @intFromFloat(@log10(@as(f64, @floatFromInt(num))) + 1));
            if (@mod(digits, 2) == 0) {
                const tens = std.math.pow(u64, 10, digits / 2);

                const num1 = num / tens;
                if (stones2.get(num1)) |c| {
                    try stones2.put(num1, c + cnt);
                } else {
                    try stones2.put(num1, cnt);
                }

                const num2 = @mod(num, tens);
                if (stones2.get(num2)) |c| {
                    try stones2.put(num2, c + cnt);
                } else {
                    try stones2.put(num2, cnt);
                }

                continue;
            }

            const n = num * 2024;
            if (stones2.get(n)) |c| {
                try stones2.put(n, c + cnt);
            } else {
                try stones2.put(n, cnt);
            }
        }

        const tmp = stones;
        stones = stones2;
        stones2 = tmp;
        stones2.clearRetainingCapacity();
    }

    var count: u64 = 0;
    var itr = stones.iterator();
    while (itr.next()) |kv| {
        const cnt = kv.value_ptr.*;
        count += cnt;
    }
    return count;
}
