# Ruby 2.6.1
3.times do
	print "ha"
end

# call below method with laugh and add optional argument. If no argument provided, will laugh 3 times
# laugh 6 => "hahahahahaha", laugh 4 => "hahahaha", laugh => "hahaha"
def laugh(funny_scale=nil)
	if funny_scale.is_a? Integer
		funny_scale.times do
			giggle
		end
	else
		chuckle
	end
end

def chuckle
	3.times do
		giggle
	end
end

def giggle
	print "ha"
end