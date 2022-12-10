cycle = 0
register = 1

sum = 0

# Open the input file in read-only mode
File.open("input", "r") do |file|
  # Loop over each line in the file
  file.each_line do |line|
    parts = line.split(" ")

    # noop just increases cycles by 1 and moves on
    if parts[0] == "noop" then
      cycle += 1

      if cycle % 40 == 20 then
        sum += (cycle * register)
      end

      next
    end

    # addx rests for two cycles, THEN adds value to register
    i = 0
    until i == 2
      cycle += 1

      if cycle % 40 == 20 then
        sum += (cycle * register)
      end

      i += 1
    end

    register += parts[1].to_i
  end
end

puts sum
