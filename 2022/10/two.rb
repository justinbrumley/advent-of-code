$cycle = 0
$register = 1

$row = ""

# Add character to CRT based on register (sprite) position
def addToCRT()
  position = $row.length
  if $register >= position - 1 and $register <= position + 1 then
    # Sprite is visible
    $row += "█"
  else
    # Sprite is NOT visible
    $row += " "
  end
end

# Print current row and clear it
def printRow()
  puts $row
  $row = ""
end

# Open the input file in read-only mode
File.open("input", "r") do |file|
  # Loop over each line in the file
  file.each_line do |line|
    parts = line.split(" ")
    cmd = parts[0]
    cycles = cmd == "noop" ? 1 : 2

    i = 0
    until i == cycles
      $cycle += 1
      addToCRT()

      if $cycle > 0 and $cycle % 40 == 0 then
        printRow()
      end

      i += 1
    end

    if cmd == "addx" then
      $register += parts[1].to_i
    end
  end
end

puts $row
