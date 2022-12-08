local function getInput ()
  file = io.open("./input")

  local input = {}
  for row in file:lines() do
    local line = {}
    row:gsub(".", function (c) line[#line+1] = tonumber(c) end)
    input[#input+1] = line
  end

  return input
end

-- Check if tree is visible from any (cardinal) direction
local function isVisible (input, x, y)
  -- Edges are always visible
  if x == 1 or y == 1 or x == #input[y] or y == #input then
    return true
  end

  local height = input[y][x]

  -- Check North/South first
  local visible = true

  for i = 1, #input, 1 do
    if i == y then
      if visible then
        break
      end

      -- Reset back to true to check other direction
      visible = true
    elseif input[i][x] >= input[y][x] then
      -- Tree is taller, thus blocking view of target tree
      visible = false
    end
  end

  if visible then
    return true
  end

  -- Reset back to visible to check East/West
  visible = true

  for i = 1, #input[y], 1 do
    if i == x then
      if visible then
        break
      end

      -- Reset back to true to check other direction
      visible = true
    elseif input[y][i] >= input[y][x] then
      -- Tree is taller, thus blocking view of target tree
      visible = false
    end
  end

  return visible
end

local function countVisibleTrees (input)
  local count = 0

  for y = 1, #input, 1 do
    for x = 1, #input[y], 1 do
      if isVisible(input, x, y) then
        count = count + 1
      end
    end
  end

  return count
end

local function getScenicScore (input, x, y)
  -- Edges have a score of 0
  if x == 1 or y == 1 or x == #input[y] or y == #input then
    return 0
  end

  local up = 0
  for i = y - 1, 1, -1 do
    up = up + 1

    if input[i][x] >= input[y][x] then
      break
    end
  end

  local down = 0
  for i = y + 1, #input, 1 do
    down = down + 1

    if input[i][x] >= input[y][x] then
      break
    end
  end

  local left = 0
  for i = x - 1, 1, -1 do
    left = left + 1

    if input[y][i] >= input[y][x] then
      break
    end
  end

  local right = 0
  for i = x + 1, #input[y], 1 do
    right = right + 1

    if input[y][i] >= input[y][x] then
      break
    end
  end

  return up * down * left * right
end

local function getBestScenicScore (input)
  local max = 0

  for y = 1, #input, 1 do
    for x = 1, #input[y], 1 do
      local score = getScenicScore(input, x, y)
      if score > max then
        max = score
      end
    end
  end

  return max
end

local input = getInput()

-- Part One (1763)
print('Visible Trees', countVisibleTrees(input))

-- Part Two (671160)
print('Best Scenic Score', getBestScenicScore(input))
