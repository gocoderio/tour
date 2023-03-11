<div>
    <ul>
        <li>
            <span>Using the tour</span>
            <ul>
                <li>
                    <span>Welcome!</span>
                    <ul>
                        <li >
                            <a id="home/src/github.com/gocoderpro/tour/_01_welcome/_01_hello" onclick="lessonOpen(this.id)">Hello, 世界</a>
                        </li>
						<li>
                            <a id="home/src/github.com/gocoderpro/tour/_01_welcome/_02_webassembly" onclick="lessonOpen(this.id)" class="greyedOut" title="Complete the prior lesson to continue here" >Go local</a>
                        </li>
						<li>
                            <a id="home/src/github.com/gocoderpro/tour/_01_welcome/_03_congratulations" onclick="lessonOpen(this.id)" class="greyedOut" title="Complete the prior lesson to continue here" lastLesson>Congratulations</a>
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
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_01_packages" onclick="lessonOpen(this.id)" class="greyedOut">Packages</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_02_imports" onclick="lessonOpen(this.id)" class="greyedOut">Imports</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_03_exported-names" onclick="lessonOpen(this.id)" class="greyedOut">Exported names</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_04_functions" onclick="lessonOpen(this.id)" class="greyedOut">Functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_05_functions-continued" onclick="lessonOpen(this.id)" class="greyedOut">Functions continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_06_multiple-results" onclick="lessonOpen(this.id)" class="greyedOut">Multiple results</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_07_named-results" onclick="lessonOpen(this.id)" class="greyedOut">Named return values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_08_variables" onclick="lessonOpen(this.id)" class="greyedOut">Variables</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_09_variables-with-initializers" onclick="lessonOpen(this.id)" class="greyedOut">Variables with initializers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope active" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_10_short-variable-declarations" onclick="lessonOpen(this.id)" class="greyedOut">Short variables declarations</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_11_basic-types" onclick="lessonOpen(this.id)" class="greyedOut">Basic types</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_12_zero" onclick="lessonOpen(this.id)" class="greyedOut">Zero values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_13_type-conversions" onclick="lessonOpen(this.id)" class="greyedOut">Type conversions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_14_type-inference" onclick="lessonOpen(this.id)" class="greyedOut">Type inference</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_15_constants" onclick="lessonOpen(this.id)" class="greyedOut">Constants</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/gocoderpro/tour/_02_basics/_16_numeric-constants" onclick="lessonOpen(this.id)" class="greyedOut">numeric constants</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-flowcontrol" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Flow control statements: for, if, else, switch and defer</span>
                    <ul>
                       <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_01_for" onclick="lessonOpen(this.id)" class="greyedOut">For</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_02_for-continued" onclick="lessonOpen(this.id)" class="greyedOut">For continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_03_for-is-gos-while" onclick="lessonOpen(this.id)" class="greyedOut">For is Go's "while"</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_04_forever" onclick="lessonOpen(this.id)" class="greyedOut">Forever</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_05_if" onclick="lessonOpen(this.id)" class="greyedOut">If</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_06_if-with-a-short-statement" onclick="lessonOpen(this.id)" class="greyedOut">If with a short statement</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_07_if-and-else" onclick="lessonOpen(this.id)" class="greyedOut">If and else</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_08_exercise-loops-and-functions" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: loops and functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_09_switch" onclick="lessonOpen(this.id)" class="greyedOut">Switch</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_10_switch-evaluation-order" onclick="lessonOpen(this.id)" class="greyedOut">Switch evaluation order</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_11_switch-with-no-condition" onclick="lessonOpen(this.id)" class="greyedOut">Switch with no condition</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_12_defer" onclick="lessonOpen(this.id)" class="greyedOut">Defer</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_03_flowcontrol/_13_defer-multi" onclick="lessonOpen(this.id)" class="greyedOut">Stacking defers</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-moretypes" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">More types: structs, slices, and maps.</span>
                    <ul>
                        <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_01_pointers" onclick="lessonOpen(this.id)" class="greyedOut">Pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_02_structs" onclick="lessonOpen(this.id)" class="greyedOut">Structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_03_struct-fields" onclick="lessonOpen(this.id)" class="greyedOut">Struct Fields</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_04_struct-pointers" onclick="lessonOpen(this.id)" class="greyedOut">Pointers to structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_05_struct-literals" onclick="lessonOpen(this.id)" class="greyedOut">Struct Literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_06_arrays" onclick="lessonOpen(this.id)" class="greyedOut">Arrays</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_07_slices" onclick="lessonOpen(this.id)" class="greyedOut">Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_08_slices-pointers" onclick="lessonOpen(this.id)" class="greyedOut">Slice pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_09_slice-literals" onclick="lessonOpen(this.id)" class="greyedOut">Slice literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_10_slice-bounds" onclick="lessonOpen(this.id)" class="greyedOut">Slice bounds</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_11_slice-len-cap" onclick="lessonOpen(this.id)" class="greyedOut">Slice length and capacity</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_12_nil-slices" onclick="lessonOpen(this.id)" class="greyedOut">Nil slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_013_making-slices" onclick="lessonOpen(this.id)" class="greyedOut">Making slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_14_slices-of-slice" onclick="lessonOpen(this.id)" class="greyedOut">Slices of slice</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_15_append" onclick="lessonOpen(this.id)" class="greyedOut">Append</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_16_range" onclick="lessonOpen(this.id)" class="greyedOut">Range</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_17_range-continued" onclick="lessonOpen(this.id)" class="greyedOut">Range continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_18_exercise-slices" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_19_maps" onclick="lessonOpen(this.id)" class="greyedOut">Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_20_map-literals" onclick="lessonOpen(this.id)" class="greyedOut">Map literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_21_map-literals-continued" onclick="lessonOpen(this.id)" class="greyedOut">Map literals continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_22_mutating-maps" onclick="lessonOpen(this.id)" class="greyedOut">Mutating maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_23_exercise-maps" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_24_function-values" onclick="lessonOpen(this.id)" class="greyedOut">Function values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_25_function closures" onclick="lessonOpen(this.id)" class="greyedOut">Function closures</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_04_moretypes/_26_exercise-fibonacci-closure" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Fibonacci closure</a>
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
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_01_methods" onclick="lessonOpen(this.id)" class="greyedOut">Methods</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_02_methods-funcs" onclick="lessonOpen(this.id)" class="greyedOut">Methods and functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_03_methods-continued" onclick="lessonOpen(this.id)" class="greyedOut">Methods continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_04_methods-pointers" onclick="lessonOpen(this.id)" class="greyedOut">Methods pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_05_methods-pointers-explained" onclick="lessonOpen(this.id)" class="greyedOut">Methods pointers explained</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_06_indirection" onclick="lessonOpen(this.id)" class="greyedOut">Inderection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_07_indirection-values" onclick="lessonOpen(this.id)" class="greyedOut">Indirection values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_08_methods-with-pointer-receivers" onclick="lessonOpen(this.id)" class="greyedOut">Methods with pointer receivers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_09_interfaces" onclick="lessonOpen(this.id)" class="greyedOut">Interfaces</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_10_interfaces-are-satisfied-implicitly" onclick="lessonOpen(this.id)" class="greyedOut">Interfaces are satisfided implicitly</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_11_interface-values" onclick="lessonOpen(this.id)" class="greyedOut">Interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_12_interface-values-with-nil" onclick="lessonOpen(this.id)" class="greyedOut">Interface values with nil</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_13_nil-interface-values" onclick="lessonOpen(this.id)" class="greyedOut">Interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_14_empty-interface" onclick="lessonOpen(this.id)" class="greyedOut">Empty Interface</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_15_type-assertions" onclick="lessonOpen(this.id)" class="greyedOut">Type assertions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_16_type-switches" onclick="lessonOpen(this.id)" class="greyedOut">Type switches</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_17_stringers" onclick="lessonOpen(this.id)" class="greyedOut">Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_18_exercise-stringers" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_19_errors" onclick="lessonOpen(this.id)" class="greyedOut">Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_20_exercise-errors" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_21_reader" onclick="lessonOpen(this.id)" class="greyedOut">Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_22_exercise-reader" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_23_exercise-rot-reader" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Rot Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_24_images" onclick="lessonOpen(this.id)" class="greyedOut">Images</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_05_methods/_25_exercise-images" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Images</a>
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
                            <a id="home/src/github.com/gocoderpro/tour/_06_generics/_01_index" onclick="lessonOpen(this.id)" class="greyedOut">Index</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_06_generics/_02_list" onclick="lessonOpen(this.id)" class="greyedOut">list</a>
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
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_01_goroutines" onclick="lessonOpen(this.id)" class="greyedOut">Goruotines</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_02_channels" onclick="lessonOpen(this.id)" class="greyedOut">Channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_03_buffered-channels" onclick="lessonOpen(this.id)" class="greyedOut">Buffered channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_04_range-and-close" onclick="lessonOpen(this.id)" class="greyedOut">Channels</a>
                            <a href="/concurrency/4" ng-click="hideTOC(true)" class="ng-binding">Range and Close</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_05_select" onclick="lessonOpen(this.id)" class="greyedOut">Select</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_06_default-selection" onclick="lessonOpen(this.id)" class="greyedOut">Default selection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_07_exercise-equivalent-binary-trees" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Equivalent binary trees</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_08_mutex-counter" onclick="lessonOpen(this.id)" class="greyedOut">Mutex counter</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/gocoderpro/tour/_07_concurrency/_09_exercise-web-crawler" onclick="lessonOpen(this.id)" class="greyedOut">Exercise: Web crawler</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
    <div class="click-catcher" ng-click="hideTOC(false)"></div>
</div>
<br><br><br><br><br>
