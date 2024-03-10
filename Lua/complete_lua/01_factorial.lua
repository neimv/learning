-- Recursive factorial
function factorial(n)
    if n == 0 then -- base case
        return 1
    else
        return n * factorial(n-1) -- recursive case
    end
end

