<h2 id=\"part2\">--- Part Two ---</h2><p>Considering every single measurement isn't as useful as you expected: there's just too much noise in the data.</p>
<p>Instead, consider sums of a <em>three-measurement sliding window</em>.  Again considering the above example:</p>
<pre><code>199  A      
200  A B    
208  A B C  
210    B C D
200  E   C D
207  E F   D
240  E F G  
269    F G H
260      G H
263        H
</code></pre>
<p>Start by comparing the first and second three-measurement windows. The measurements in the first window are marked <code>A</code> (<code>199</code>, <code>200</code>, <code>208</code>); their sum is <code>199 + 200 + 208 = 607</code>. The second window is marked <code>B</code> (<code>200</code>, <code>208</code>, <code>210</code>); its sum is <code>618</code>. The sum of measurements in the second window is larger than the sum of the first, so this first comparison <em>increased</em>.</p>
<p>Your goal now is to count <em>the number of times the sum of measurements in this sliding window increases</em> from the previous sum. So, compare <code>A</code> with <code>B</code>, then compare <code>B</code> with <code>C</code>, then <code>C</code> with <code>D</code>, and so on. Stop when there aren't enough measurements left to create a new three-measurement sum.</p>
<p>In the above example, the sum of each three-measurement window is as follows:</p>
<pre><code>A: 607 (N/A - no previous sum)
B: 618 (<em>increased</em>)
C: 618 (no change)
D: 617 (decreased)
E: 647 (<em>increased</em>)
F: 716 (<em>increased</em>)
G: 769 (<em>increased</em>)
H: 792 (<em>increased</em>)
</code></pre>
<p>In this example, there are <em><code>5</code></em> sums that are larger than the previous sum.</p>
<p>Consider sums of a three-measurement sliding window. <em>How many sums are larger than the previous sum?</em></p>
