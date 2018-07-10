class HashTable
  attr_accessor :table

  def initialize
    @table = Array.new(10) { |x| x = LinkedList.new }
  end

  # what is that key for?
  def put(_key, value)
    value = value.to_s
    index = value.size

    @table[index].insert(value)
  end

  def get(value)
    index = value.size

    @table[index].find(value)
  end

  def tprint 
    @table.each.with_index do |item, index|
      print "#{index} -> "
      item.print_list
      print "\n"
    end
  end
end

class LinkedList 
  attr_accessor :data, :next
  
  def initialize(data = nil)
    @data = data
    @next = nil
  end

  def print_list
    if @data
      print "#{@data}, "
    end

    if @next
      @next.print_list
    end
  end

  def insert(value)
    if @data.nil?
      @data, @next = value, LinkedList.new
    else
      return @next.insert(value)
    end

    return self
  end

  def find(value)
    if @data.nil?
      return "Nonexistant"
    else
      @data == value ? data : @next.find(value)
    end
  end
end
