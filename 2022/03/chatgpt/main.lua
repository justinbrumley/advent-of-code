-- Open the file for reading
local file = io.open("input.txt", "r")

-- Create a variable to store the total score
local totalScore = 0

-- Iterate through each line in the file
for line in file:lines() do
  -- Split the line in half
  local firstHalf = line:sub(1, math.floor(line:len() / 2))
  local secondHalf = line:sub(math.floor(line:len() / 2) + 1)

  -- Create a variable to store the first shared character
  local firstSharedChar = ""

  -- Iterate through each character in the first half of the line
  for i = 1, firstHalf:len() do
    -- Get the character at the current index
    local char = firstHalf:sub(i, i)

    -- Check if the character is also in the second half of the line
    if secondHalf:find(char) then
      -- If so, store the character as the first shared character
      firstSharedChar = char
      break
    end
  end

  -- Score the first shared character
  local score = 0
  if firstSharedChar:find("[a-zA-Z]") then
    -- If the character is a letter, add the corresponding score to the total
    if firstSharedChar:find("[a-z]") then
      score = (string.byte(firstSharedChar) - 96)
    else
      score = (string.byte(firstSharedChar) - 64 + 26)
    end
  end

  -- Add the score for this line to the total score
  totalScore = totalScore + score
end

-- Print the total score to the console
print(totalScore)

-- Close the file
file:close()

