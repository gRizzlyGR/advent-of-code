<h2 id=\"part2\">--- Part Two ---</h2><p>Next, you should verify the <em>life support rating</em>, which can be determined by multiplying the <em>oxygen generator rating</em> by the <em>CO2 scrubber rating</em>.</p>
<p>Both the oxygen generator rating and the CO2 scrubber rating are values that can be found in your diagnostic report - finding them is the tricky part. Both values are located using a similar process that involves filtering out values until only one remains. Before searching for either rating value, start with the full list of binary numbers from your diagnostic report and <em>consider just the first bit</em> of those numbers. Then:</p>
<ul>
<li>Keep only numbers selected by the <em>bit criteria</em> for the type of rating value for which you are searching. Discard numbers which do not match the bit criteria.</li>
<li>If you only have one number left, stop; this is the rating value for which you are searching.</li>
<li>Otherwise, repeat the process, considering the next bit to the right.</li>
</ul>
<p>The <em>bit criteria</em> depends on which type of rating value you want to find:</p>
<ul>
<li>To find <em>oxygen generator rating</em>, determine the <em>most common</em> value (<code>0</code> or <code>1</code>) in the current bit position, and keep only numbers with that bit in that position. If <code>0</code> and <code>1</code> are equally common, keep values with a <code><em>1</em></code> in the position being considered.</li>
<li>To find <em>CO2 scrubber rating</em>, determine the <em>least common</em> value (<code>0</code> or <code>1</code>) in the current bit position, and keep only numbers with that bit in that position. If <code>0</code> and <code>1</code> are equally common, keep values with a <code><em>0</em></code> in the position being considered.</li>
</ul>
<p>For example, to determine the <em>oxygen generator rating</em> value using the same example diagnostic report from above:</p>
<ul>
<li>Start with all 12 numbers and consider only the first bit of each number. There are more <code>1</code> bits (7) than <code>0</code> bits (5), so keep only the 7 numbers with a <code>1</code> in the first position: <code>11110</code>, <code>10110</code>, <code>10111</code>, <code>10101</code>, <code>11100</code>, <code>10000</code>, and <code>11001</code>.</li>
<li>Then, consider the second bit of the 7 remaining numbers: there are more <code>0</code> bits (4) than <code>1</code> bits (3), so keep only the 4 numbers with a <code>0</code> in the second position: <code>10110</code>, <code>10111</code>, <code>10101</code>, and <code>10000</code>.</li>
<li>In the third position, three of the four numbers have a <code>1</code>, so keep those three: <code>10110</code>, <code>10111</code>, and <code>10101</code>.</li>
<li>In the fourth position, two of the three numbers have a <code>1</code>, so keep those two: <code>10110</code> and <code>10111</code>.</li>
<li>In the fifth position, there are an equal number of <code>0</code> bits and <code>1</code> bits (one each). So, to find the <em>oxygen generator rating</em>, keep the number with a <code>1</code> in that position: <code>10111</code>.</li>
<li>As there is only one number left, stop; the <em>oxygen generator rating</em> is <code>10111</code>, or <code><em>23</em></code> in decimal.</li>
</ul>
<p>Then, to determine the <em>CO2 scrubber rating</em> value from the same example above:</p>
<ul>
<li>Start again with all 12 numbers and consider only the first bit of each number. There are fewer <code>0</code> bits (5) than <code>1</code> bits (7), so keep only the 5 numbers with a <code>0</code> in the first position: <code>00100</code>, <code>01111</code>, <code>00111</code>, <code>00010</code>, and <code>01010</code>.</li>
<li>Then, consider the second bit of the 5 remaining numbers: there are fewer <code>1</code> bits (2) than <code>0</code> bits (3), so keep only the 2 numbers with a <code>1</code> in the second position: <code>01111</code> and <code>01010</code>.</li>
<li>In the third position, there are an equal number of <code>0</code> bits and <code>1</code> bits (one each). So, to find the <em>CO2 scrubber rating</em>, keep the number with a <code>0</code> in that position: <code>01010</code>.</li>
<li>As there is only one number left, stop; the <em>CO2 scrubber rating</em> is <code>01010</code>, or <code><em>10</em></code> in decimal.</li>
</ul>
<p>Finally, to find the life support rating, multiply the oxygen generator rating (<code>23</code>) by the CO2 scrubber rating (<code>10</code>) to get <code><em>230</em></code>.</p>
<p>Use the binary numbers in your diagnostic report to calculate the oxygen generator rating and CO2 scrubber rating, then multiply them together. <em>What is the life support rating of the submarine?</em> (Be sure to represent your answer in decimal, not binary.)</p>
