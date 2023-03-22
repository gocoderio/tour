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
                        <li >
                            <a id="home/src/github.com/gocoderio/tour/_01_welcome/_01_hello" onclick="lessonOpen(event)">Hello, 世界</a>
                        </li>
						<li>
                            <a id="home/src/github.com/gocoderio/tour/_01_welcome/_02_webassembly" onclick="lessonOpen(event)"  title="Complete the prior lesson to continue here">Go local</a>
                        </li>
						<li>
                            <a id="home/src/github.com/gocoderio/tour/_01_welcome/_03_congratulations" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li ng-repeat="m in toc.modules" class="toc-module ng-scope" id="toc-m-basics">
            <span class="ng-binding">Basics</span>
            <ul>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope active" id="toc-l-basics" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Packages, variables, and functions.</span>
                    <ul>
                        <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_01_packages" onclick="lessonOpen(event)">Packages</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_02_imports" onclick="lessonOpen(event)">Imports</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_03_exported-names" onclick="lessonOpen(event)">Exported names</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_04_functions" onclick="lessonOpen(event)">Functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_05_functions-continued" onclick="lessonOpen(event)">Functions continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_06_multiple-results" onclick="lessonOpen(event)">Multiple results</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_07_named-results" onclick="lessonOpen(event)">Named return values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_08_variables" onclick="lessonOpen(event)">Variables</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_09_variables-with-initializers" onclick="lessonOpen(event)">Variables with initializers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope active" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_10_short-variable-declarations" onclick="lessonOpen(event)">Short variables declarations</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_11_basic-types" onclick="lessonOpen(event)">Basic types</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_12_zero" onclick="lessonOpen(event)">Zero values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_13_type-conversions" onclick="lessonOpen(event)">Type conversions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_14_type-inference" onclick="lessonOpen(event)">Type inference</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_15_constants" onclick="lessonOpen(event)">Constants</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_16_numeric-constants" onclick="lessonOpen(event)">numeric constants</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_02_basics/_17_congratulations" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-flowcontrol" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Flow control statements: for, if, else, switch and defer</span>
                    <ul>
                       <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_01_for" onclick="lessonOpen(event)" class="greyedOut">For</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_02_for-continued" onclick="lessonOpen(event)" class="greyedOut">For continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_03_for-is-gos-while" onclick="lessonOpen(event)" class="greyedOut">For is Go's "while"</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_04_forever" onclick="lessonOpen(event)" class="greyedOut">Forever</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_05_if" onclick="lessonOpen(event)" class="greyedOut">If</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_06_if-with-a-short-statement" onclick="lessonOpen(event)" class="greyedOut">If with a short statement</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_07_if-and-else" onclick="lessonOpen(event)" class="greyedOut">If and else</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_08_exercise-loops-and-functions" onclick="lessonOpen(event)" class="greyedOut">Exercise: loops and functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_09_switch" onclick="lessonOpen(event)" class="greyedOut">Switch</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_10_switch-evaluation-order" onclick="lessonOpen(event)" class="greyedOut">Switch evaluation order</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_11_switch-with-no-condition" onclick="lessonOpen(event)" class="greyedOut">Switch with no condition</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_12_defer" onclick="lessonOpen(event)" class="greyedOut">Defer</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_13_defer-multi" onclick="lessonOpen(event)" class="greyedOut">Stacking defers</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_03_flowcontrol/_14_congratulations" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-moretypes" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">More types: structs, slices, and maps.</span>
                    <ul>
                        <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_01_pointers" onclick="lessonOpen(event)" class="greyedOut">Pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_02_structs" onclick="lessonOpen(event)" class="greyedOut">Structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_03_struct-fields" onclick="lessonOpen(event)" class="greyedOut">Struct Fields</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_04_struct-pointers" onclick="lessonOpen(event)" class="greyedOut">Pointers to structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_05_struct-literals" onclick="lessonOpen(event)" class="greyedOut">Struct Literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_06_array" onclick="lessonOpen(event)" class="greyedOut">Arrays</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_07_slices" onclick="lessonOpen(event)" class="greyedOut">Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_08_slices-pointers" onclick="lessonOpen(event)" class="greyedOut">Slice pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_09_slice-literals" onclick="lessonOpen(event)" class="greyedOut">Slice literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_10_slice-bounds" onclick="lessonOpen(event)" class="greyedOut">Slice bounds</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_11_slice-len-cap" onclick="lessonOpen(event)" class="greyedOut">Slice length and capacity</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_12_nil-slices" onclick="lessonOpen(event)" class="greyedOut">Nil slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_13_making-slices" onclick="lessonOpen(event)" class="greyedOut">Making slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_14_slices-of-slice" onclick="lessonOpen(event)" class="greyedOut">Slices of slice</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_15_append" onclick="lessonOpen(event)" class="greyedOut">Append</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_16_range" onclick="lessonOpen(event)" class="greyedOut">Range</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_17_range-continued" onclick="lessonOpen(event)" class="greyedOut">Range continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_18_exercise-slices" onclick="lessonOpen(event)" class="greyedOut">Exercise: Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_19_maps" onclick="lessonOpen(event)" class="greyedOut">Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_20_map-literals" onclick="lessonOpen(event)" class="greyedOut">Map literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_21_map-literals-continued" onclick="lessonOpen(event)" class="greyedOut">Map literals continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_22_mutating-maps" onclick="lessonOpen(event)" class="greyedOut">Mutating maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_23_exercise-maps" onclick="lessonOpen(event)" class="greyedOut">Exercise: Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_24_function-values" onclick="lessonOpen(event)" class="greyedOut">Function values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_25_function-closures" onclick="lessonOpen(event)" class="greyedOut">Function closures</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_26_exercise-fibonacci-closure" onclick="lessonOpen(event)" class="greyedOut">Exercise: Fibonacci closure</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_04_moretypes/_27_congratulations" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li ng-repeat="m in toc.modules" class="toc-module ng-scope" id="toc-m-methods">
            <span class="ng-binding">Methods and interfaces</span>
            <ul>
                <!-- ngRepeat: l in m.lessons --><li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-methods" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Methods and interfaces</span>
                    <ul>
                        <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_01_methods" onclick="lessonOpen(event)" class="greyedOut">Methods</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_02_methods-funcs" onclick="lessonOpen(event)" class="greyedOut">Methods and functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_03_methods-continued" onclick="lessonOpen(event)" class="greyedOut">Methods continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_04_methods-pointers" onclick="lessonOpen(event)" class="greyedOut">Methods pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_05_methods-pointers-explained" onclick="lessonOpen(event)" class="greyedOut">Methods pointers explained</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_06_indirection" onclick="lessonOpen(event)" class="greyedOut">Inderection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_07_indirection-values" onclick="lessonOpen(event)" class="greyedOut">Indirection values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_08_methods-with-pointer-receivers" onclick="lessonOpen(event)" class="greyedOut">Methods with pointer receivers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_09_interfaces" onclick="lessonOpen(event)" class="greyedOut">Interfaces</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_10_interfaces-are-satisfied-implicitly" onclick="lessonOpen(event)" class="greyedOut">Interfaces are satisfided implicitly</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_11_interface-values" onclick="lessonOpen(event)" class="greyedOut">Interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_12_interface-values-with-nil" onclick="lessonOpen(event)" class="greyedOut">Interface values with nil</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_13_nil-interface-values" onclick="lessonOpen(event)" class="greyedOut">Interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_14_empty-interface" onclick="lessonOpen(event)" class="greyedOut">Empty Interface</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_15_type-assertions" onclick="lessonOpen(event)" class="greyedOut">Type assertions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_16_type-switches" onclick="lessonOpen(event)" class="greyedOut">Type switches</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_17_stringer" onclick="lessonOpen(event)" class="greyedOut">Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_18_exercise-stringer" onclick="lessonOpen(event)" class="greyedOut">Exercise: Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_19_errors" onclick="lessonOpen(event)" class="greyedOut">Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_20_exercise-errors" onclick="lessonOpen(event)" class="greyedOut">Exercise: Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_21_reader" onclick="lessonOpen(event)" class="greyedOut">Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_22_exercise-reader" onclick="lessonOpen(event)" class="greyedOut">Exercise: Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_23_exercise-rot-reader" onclick="lessonOpen(event)" class="greyedOut">Exercise: Rot Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_24_images" onclick="lessonOpen(event)" class="greyedOut">Images</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_25_exercise-images" onclick="lessonOpen(event)" class="greyedOut">Exercise: Images</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_05_methods/_26_congratulations" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li ng-repeat="m in toc.modules" class="toc-module ng-scope" id="toc-m-generics">
            <span class="ng-binding">Generics</span>
            <ul>
                <!-- ngRepeat: l in m.lessons --><li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-generics" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Generics</span>
                    <ul>
                        <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_06_generics/_01_index" onclick="lessonOpen(event)" class="greyedOut">Index</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_06_generics/_02_list" onclick="lessonOpen(event)" class="greyedOut">list</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_06_generics/_03_congratulations" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li><li ng-repeat="m in toc.modules" class="toc-module ng-scope" id="toc-m-concurrency">
            <span class="ng-binding">Concurrency</span>
            <ul>
                <!-- ngRepeat: l in m.lessons --><li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-concurrency" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Concurrency</span>
                    <ul>
                        <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_01_goroutines" onclick="lessonOpen(event)" class="greyedOut">Goroutines</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_02_channels" onclick="lessonOpen(event)" class="greyedOut">Channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_03_buffered-channels" onclick="lessonOpen(event)" class="greyedOut">Buffered channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_04_range-and-close" onclick="lessonOpen(event)" class="greyedOut">Range and Close</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_05_select" onclick="lessonOpen(event)" class="greyedOut">Select</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_06_default-selection" onclick="lessonOpen(event)" class="greyedOut">Default selection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_07_exercise-equivalent-binary-trees" onclick="lessonOpen(event)" class="greyedOut">Exercise: Equivalent binary trees</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_08_mutex-counter" onclick="lessonOpen(event)" class="greyedOut">Mutex counter</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_09_exercise-web-crawler" onclick="lessonOpen(event)" class="greyedOut">Exercise: Web crawler</a>
                        </li>
                        <li>
                            <a id="home/src/github.com/gocoderio/tour/_07_concurrency/_10_congratulations" onclick="lessonOpen(event)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
    <div class="click-catcher" ng-click="hideTOC(false)"></div>
</div>
<br><br><br><br><br>
