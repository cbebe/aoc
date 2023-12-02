const std = @import("std");

pub fn main() !void {}

// const data_2a = .{ @embedFile("./input/2/test.txt"), 8 };
// const data_2b = .{ @embedFile("./input/2/test.txt"), 2286 };
const data_2a = .{ @embedFile("./input/2/input.txt"), 2439 };
const data_2b = .{ @embedFile("./input/2/input.txt"), 63711 };

pub fn split(x: []const u8, c: []const u8) std.mem.SplitIterator(u8) {
    return std.mem.split(u8, x, c);
}

pub fn trim_char(x: []const u8, c: []const u8) []const u8 {
    return std.mem.trim(u8, x, c);
}

pub fn trim(x: []const u8) []const u8 {
    return std.mem.trim(u8, x, " ");
}

pub fn int(x: []const u8) !i32 {
    return std.fmt.parseInt(i32, x, 10);
}

pub fn solve_2a(d: []const u8, alloc: std.mem.Allocator) !i32 {
    var it = split(d, "\n");
    var sum: i32 = 0;
    var map = std.StringHashMap(i32).init(alloc);
    defer map.deinit();
    try map.put("green", 13);
    try map.put("red", 12);
    try map.put("blue", 14);
    while (it.next()) |x| {
        if (x.len == 0) {
            continue;
        }
        var fail = false;
        var it2 = split(x, ":");
        const a = it2.next().?;
        const b = trim(it2.next().?);
        var it3 = split(a, " ");
        _ = it3.next();
        const game_id = try int(it3.next().?);

        var sets = split(b, ";");
        while (sets.next()) |s| {
            var cubes = split(trim(s), ", ");
            while (cubes.next()) |c| {
                var t = split(c, " ");
                const val = try int(t.next().?);
                if (val > map.get(t.next().?).?) {
                    fail = true;
                    break;
                }
            }
            if (fail) {
                break;
            }
        }
        if (!fail) {
            sum += game_id;
        }
    }

    return sum;
}

pub fn solve_2b(d: []const u8, alloc: std.mem.Allocator) !i32 {
    var it = split(d, "\n");
    var sum: i32 = 0;
    var map = std.StringHashMap(i32).init(alloc);
    defer map.deinit();
    while (it.next()) |x| {
        if (x.len == 0) {
            continue;
        }
        try map.put("green", 0);
        try map.put("red", 0);
        try map.put("blue", 0);
        var it2 = split(x, ":");
        const a = it2.next().?;
        const b = trim(it2.next().?);
        var it3 = split(a, " ");
        _ = it3.next();
        const game_id = try int(it3.next().?);
        _ = game_id;

        var sets = split(b, ";");
        while (sets.next()) |s| {
            var cubes = split(trim(s), ", ");
            while (cubes.next()) |c| {
                var t = split(c, " ");
                const val = try int(t.next().?);
                const colour = t.next().?;
                if (val > map.get(colour).?) {
                    try map.put(colour, val);
                }
            }
        }
        sum += map.get("red").? *
            map.get("green").? *
            map.get("blue").?;
    }

    return sum;
}

test "day2" {
    try std.testing.expectEqual(@as(i32, data_2a[1]), try solve_2a(data_2a[0], std.testing.allocator));
    try std.testing.expectEqual(@as(i32, data_2b[1]), try solve_2b(data_2b[0], std.testing.allocator));
}
