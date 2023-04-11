<div id="menu">
<!--div id="menu" style="height:300px;overflow-y:auto"-->
    <!--menu id tag needed for forward and backward buttons to work through this menu-->
    <ul>
        <li>
            <span>Using the tour</span>
            <ul>
                <li>
                    <span>Welcome!</span>
                    <ul>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_01_welcome/_01_hello" class="lessonLink" onclick="lessonOpen(event)">Hello, 世界</a>
                        </li>
						<li>
                            <a id="home/src/github.com/gocoderio/tour/_01_welcome/_02_webassembly" class="lessonLink" onclick="lessonOpen(event)"  title="Complete the prior lesson to continue here">Go local</a>
                        </li>
						<li>
                            <a id="home/src/github.com/gocoderio/tour/_01_welcome/_03_congratulations" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li>
            <span class="ng-binding">Basics</span>
            <ul>
                <li>
                    <span ng-click="toggleLesson(l)" class="ng-binding">Packages, variables, and functions.</span>
                    <ul>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_01_packages" class="lessonLink" onclick="lessonOpen(event)">Packages</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_02_imports" class="lessonLink" onclick="lessonOpen(event)">Imports</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_03_exported-names" class="lessonLink" onclick="lessonOpen(event)">Exported names</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_04_functions" class="lessonLink" onclick="lessonOpen(event)">Functions</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_05_functions-continued" class="lessonLink" onclick="lessonOpen(event)">Functions continued</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_06_multiple-results" class="lessonLink" onclick="lessonOpen(event)">Multiple results</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_07_named-results" class="lessonLink" onclick="lessonOpen(event)">Named return values</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_08_variables" class="lessonLink" onclick="lessonOpen(event)">Variables</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_09_variables-with-initializers" class="lessonLink" onclick="lessonOpen(event)">Variables with initializers</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_10_short-variable-declarations" class="lessonLink" onclick="lessonOpen(event)">Short variables declarations</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_11_basic-types" class="lessonLink" onclick="lessonOpen(event)">Basic types</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_12_zero" class="lessonLink" onclick="lessonOpen(event)">Zero values</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_13_type-conversions" class="lessonLink" onclick="lessonOpen(event)">Type conversions</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_14_type-inference" class="lessonLink" onclick="lessonOpen(event)">Type inference</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_15_constants" class="lessonLink" onclick="lessonOpen(event)">Constants</a>
                        </li><li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_16_numeric-constants" class="lessonLink" onclick="lessonOpen(event)">Numeric constants</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_17_congratulations" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
                <li>
                    <span ng-click="toggleLesson(l)" class="ng-binding">Flow control statements: for, if, else, switch and defer</span>
                    <ul>
                       <li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_01_for" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">For</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_02_for-continued" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">For continued</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_03_for-is-gos-while" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">For is Go's "while"</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_04_forever" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Forever</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_05_if" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">If</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_06_if-with-a-short-statement" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">If with a short statement</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_07_if-and-else" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">If and else</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_08_exercise-loops-and-functions" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: loops and functions</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_09_switch" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Switch</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_10_switch-evaluation-order" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Switch evaluation order</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_11_switch-with-no-condition" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Switch with no condition</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_12_defer" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Defer</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_13_defer-multi" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Stacking defers</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_14_congratulations" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
                <li>
                    <span ng-click="toggleLesson(l)" class="ng-binding">More types: structs, slices, and maps.</span>
                    <ul>
                        <li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_01_pointers" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Pointers</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_02_structs" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Structs</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_03_struct-fields" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Struct Fields</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_04_struct-pointers" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Pointers to structs</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_05_struct-literals" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Struct Literals</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_06_array" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Arrays</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_07_slices" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Slices</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_08_slices-pointers" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Slice pointers</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_09_slice-literals" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Slice literals</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_10_slice-bounds" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Slice bounds</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_11_slice-len-cap" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Slice length and capacity</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_12_nil-slices" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Nil slices</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_13_making-slices" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Making slices</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_14_slices-of-slice" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Slices of slice</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_15_append" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Append</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_16_range" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Range</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_17_range-continued" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Range continued</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_18_exercise-slices" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Slices</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_19_maps" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Maps</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_20_map-literals" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Map literals</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_21_map-literals-continued" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Map literals continued</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_22_mutating-maps" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Mutating maps</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_23_exercise-maps" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Maps</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_24_function-values" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Function values</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_25_function-closures" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Function closures</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_26_exercise-fibonacci-closure" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Fibonacci closure</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_27_congratulations" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li>
            <span class="ng-binding">Methods and interfaces</span>
            <ul>
                <li>
                    <span ng-click="toggleLesson(l)" class="ng-binding">Methods and interfaces</span>
                    <ul>
                        <li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_01_methods" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Methods</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_02_methods-funcs" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Methods and functions</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_03_methods-continued" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Methods continued</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_04_methods-pointers" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Methods pointers</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_05_methods-pointers-explained" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Methods pointers explained</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_06_indirection" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Inderection</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_07_indirection-values" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Indirection values</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_08_methods-with-pointer-receivers" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Methods with pointer receivers</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_09_interfaces" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Interfaces</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_10_interfaces-are-satisfied-implicitly" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Interfaces are satisfided implicitly</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_11_interface-values" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Interface values</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_12_interface-values-with-nil" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Interface values with nil</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_13_nil-interface-values" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Interface values</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_14_empty-interface" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Empty Interface</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_15_type-assertions" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Type assertions</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_16_type-switches" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Type switches</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_17_stringer" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Stringers</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_18_exercise-stringer" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Stringers</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_19_errors" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Errors</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_20_exercise-errors" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Errors</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_21_reader" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Reader</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_22_exercise-reader" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Reader</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_23_exercise-rot-reader" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Rot Reader</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_24_images" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Images</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_25_exercise-images" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Images</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_26_congratulations" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li>
            <span class="ng-binding">Generics</span>
            <ul>
                <li>
                    <span ng-click="toggleLesson(l)" class="ng-binding">Generics</span>
                    <ul>
                        <li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_06_generics/_01_index" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Index</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_06_generics/_02_list" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">list</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_06_generics/_03_congratulations" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li>
            <span class="ng-binding">Concurrency</span>
            <ul>
                <li>
                    <span ng-click="toggleLesson(l)" class="ng-binding">Concurrency</span>
                    <ul>
                        <li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_01_goroutines" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Goroutines</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_02_channels" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Channels</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_03_buffered-channels" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Buffered channels</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_04_range-and-close" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Range and Close</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_05_select" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Select</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_06_default-selection" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Default selection</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_07_exercise-equivalent-binary-trees" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Equivalent binary trees</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_08_mutex-counter" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Mutex counter</a>
                        </li><li style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_09_exercise-web-crawler" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut">Exercise: Web crawler</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_10_congratulations" class="lessonLink" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
    <div class="click-catcher" ng-click="hideTOC(false)"></div>
</div>
<br><br><br><br><br>
