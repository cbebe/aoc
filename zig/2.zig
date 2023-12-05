const std = @import("std");

pub fn main() !void {}

// const data_2a = .{ @embedFile("./input/2/test.txt"), 8 };
// const data_2b = .{ @embedFile("./input/2/test.txt"), 2286 };
const data_2a = .{ @embedFile("./input/2/input.txt"), 2439 };
const data_2b = .{ @embedFile("./input/2/input.txt"), 63711 };

const A = @import("aoc.zig");

pub fn solve2a(data: []const u8, alloc: std.mem.Allocator) !i32 {
    var it = A.split(data, "\n");
    var sum: i32 = 0;
    var map = std.StringHashMap(i32).init(alloc);
    try map.put("green", 13);
    try map.put("red", 12);
    try map.put("blue", 14);
    defer map.deinit();
    while (it.next()) |x| {
        if (x.len == 0) {
            continue;
        }
        var fail = false;
        var it2 = A.split(x, ":");
        const a = it2.next().?;
        const b = A.trim(it2.next().?);
        var it3 = A.split(a, " ");
        _ = it3.next();
        const game_id = try A.int(it3.next().?);

        var sets = A.split(b, ";");
        while (sets.next()) |s| {
            var cubes = A.split(A.trim(s), ", ");
            while (cubes.next()) |c| {
                var t = A.split(c, " ");
                const val = try A.int(t.next().?);
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

pub fn solve2b(data: []const u8, alloc: std.mem.Allocator) !i32 {
    var it = A.split(data, "\n");
    var sum: i32 = 0;
    while (it.next()) |x| {
        if (x.len == 0) {
            continue;
        }
        var map = std.StringHashMap(i32).init(alloc);
        defer map.deinit();
        try map.put("green", 0);
        try map.put("red", 0);
        try map.put("blue", 0);
        var it2 = A.split(x, ":");
        const a = it2.next().?;
        const b = A.trim(it2.next().?);
        var it3 = A.split(a, " ");
        _ = it3.next();
        const game_id = try A.int(it3.next().?);
        _ = game_id;

        var sets = A.split(b, ";");
        while (sets.next()) |s| {
            var cubes = A.split(A.trim(s), ", ");
            while (cubes.next()) |c| {
                var t = A.split(c, " ");
                const val = try A.int(t.next().?);
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
    try A.testDay(solve2a, data_2a);
    try A.testDay(solve2b, data_2b);
}
