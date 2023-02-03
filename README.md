<div>
    <ul>
        <li>
            <span>Using the tour</span>
            <ul>
                <li>
                    <span>Welcome!</span>
                    <ul>
                        <li >
                            <a id="home/src/github.com/dougwatson/tour/_01_welcome/_01_hello" onclick="pathOpen(this.id)">Hello, 世界</a>
                        </li>
						<li>
                            <a id="home/src/github.com/dougwatson/tour/_01_welcome/_02_webassembly" onclick="pathOpen(this.id)">Go local</a>
                        </li>
						<li>
                            <a id="home/src/github.com/dougwatson/tour/_01_welcome/_03_congratulations" onclick="pathOpen(this.id)">Congratulations</a>
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
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_01_packages" onclick="pathOpen(this.id)">Packages</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_02_imports" onclick="pathOpen(this.id)">Imports</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_03_exported-names" onclick="pathOpen(this.id)">Exported names</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_04_functions" onclick="pathOpen(this.id)">Functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_05_functions-continued" onclick="pathOpen(this.id)">Functions continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_06_multiple-results" onclick="pathOpen(this.id)">Multiple results</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_07_named-results" onclick="pathOpen(this.id)">Named return values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_08_variables" onclick="pathOpen(this.id)">Variables</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_09_variables-with-initializers" onclick="pathOpen(this.id)">Variables with initializers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope active" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_10_short-variable-declarations" onclick="pathOpen(this.id)">Short variables declarations</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_11_basic-types" onclick="pathOpen(this.id)">Basic types</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_12_zero" onclick="pathOpen(this.id)">Zero values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_13_type-conversions" onclick="pathOpen(this.id)">Type conversions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_14_type-inference" onclick="pathOpen(this.id)">Type inference</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_15_constants" onclick="pathOpen(this.id)">Constants</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a id="home/src/github.com/dougwatson/tour/_02_basics/_16_numeric-constants" onclick="pathOpen(this.id)">numeric constants</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-flowcontrol" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Flow control statements: for, if, else, switch and defer</span>
                    <ul>
                       <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_01_for" onclick="pathOpen(this.id)">For</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_02_for-continued" onclick="pathOpen(this.id)">For continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_03_for-is-gos-while" onclick="pathOpen(this.id)">For is Go's "while"</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_04_forever" onclick="pathOpen(this.id)">Forever</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_05_if" onclick="pathOpen(this.id)">If</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_06_if-with-a-short-statement" onclick="pathOpen(this.id)">If with a short statement</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_07_if-and-else" onclick="pathOpen(this.id)">If and else</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_08_exercise-loops-and-functions" onclick="pathOpen(this.id)">Exercise: loops and functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_09_switch" onclick="pathOpen(this.id)">Switch</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_10_switch-evaluation-order" onclick="pathOpen(this.id)">Switch evaluation order</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_11_switch-with-no-condition" onclick="pathOpen(this.id)">Switch with no condition</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_12_defer" onclick="pathOpen(this.id)">Defer</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_03_flowcontrol/_13_defer-multi" onclick="pathOpen(this.id)">Stacking defers</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-moretypes" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">More types: structs, slices, and maps.</span>
                    <ul>
                        <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_01_pointers" onclick="pathOpen(this.id)">Pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_02_structs" onclick="pathOpen(this.id)">Structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_03_struct-fields" onclick="pathOpen(this.id)">Struct Fields</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_04_pointers-to-structs" onclick="pathOpen(this.id)">Pointers to structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_05_struct-literals" onclick="pathOpen(this.id)">Struct Literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_06_arrays" onclick="pathOpen(this.id)">Arrays</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_07_slices" onclick="pathOpen(this.id)">Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_08_slices-pointers" onclick="pathOpen(this.id)">Slice pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_09_slice-literals" onclick="pathOpen(this.id)">Slice literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_10_slice-bounds" onclick="pathOpen(this.id)">Slice bounds</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_11_slice-len-cap" onclick="pathOpen(this.id)">Slice length and capacity</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_12_nil-slices" onclick="pathOpen(this.id)">Nil slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_013_making-slices" onclick="pathOpen(this.id)">Making slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_14_slices-of-slice" onclick="pathOpen(this.id)">Slices of slice</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_15_append" onclick="pathOpen(this.id)">Append</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_16_range" onclick="pathOpen(this.id)">Range</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_17_range-continued" onclick="pathOpen(this.id)">Range continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_18_exercise-slices" onclick="pathOpen(this.id)">Exercise: Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_19_maps" onclick="pathOpen(this.id)">Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_20_map-literals" onclick="pathOpen(this.id)">Map literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_21_map-literals-continued" onclick="pathOpen(this.id)">Map literals continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_22_mutating-maps" onclick="pathOpen(this.id)">Mutating maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_23_exercise-maps" onclick="pathOpen(this.id)">Exercise: Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_24_function-values" onclick="pathOpen(this.id)">Function values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_25_function closures" onclick="pathOpen(this.id)">Function closures</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_04_moretypes/_26_exercise-fibonacci-closure" onclick="pathOpen(this.id)">Exercise: Fibonacci closure</a>
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
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_01_methods" onclick="pathOpen(this.id)">Methods</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_02_methods-funcs" onclick="pathOpen(this.id)">Methods and functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_03_methods-continued" onclick="pathOpen(this.id)">Methods continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_04_methods-pointers" onclick="pathOpen(this.id)">Methods pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_05_methods-pointers-explained" onclick="pathOpen(this.id)">Methods pointers explained</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_06_indirection" onclick="pathOpen(this.id)">Inderection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_07_indirection-values" onclick="pathOpen(this.id)">Indirection values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_08_methods-with-pointer-receivers" onclick="pathOpen(this.id)">Methods with pointer receivers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_09_interfaces" onclick="pathOpen(this.id)">Interfaces</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_10_interfaces-are-satisfied-implicitly" onclick="pathOpen(this.id)">Interfaces are satisfided implicitly</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_11_interface-values" onclick="pathOpen(this.id)">Interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_12_interface-values-with-nil" onclick="pathOpen(this.id)">Interface values with nil</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_13_nil-interface-values" onclick="pathOpen(this.id)">Interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_14_empty-interface" onclick="pathOpen(this.id)">Empty Interface</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_15_type-assertions" onclick="pathOpen(this.id)">Type assertions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_16_type-switches" onclick="pathOpen(this.id)">Type switches</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_17_stringers" onclick="pathOpen(this.id)">Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_18_exercise-stringers" onclick="pathOpen(this.id)">Exercise: Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_19_errors" onclick="pathOpen(this.id)">Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_20_exercise-errors" onclick="pathOpen(this.id)">Exercise: Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_21_reader" onclick="pathOpen(this.id)">Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_22_exercise-reader" onclick="pathOpen(this.id)">Exercise: Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_23_exercise-rot-reader" onclick="pathOpen(this.id)">Exercise: Rot Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_24_images" onclick="pathOpen(this.id)">Images</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_05_methods/_25_exercise-images" onclick="pathOpen(this.id)">Exercise: Images</a>
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
                            <a id="home/src/github.com/dougwatson/tour/_06_generics/_01_index" onclick="pathOpen(this.id)">Index</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_06_generics/_02_list" onclick="pathOpen(this.id)">list</a>
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
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_01_goroutines" onclick="pathOpen(this.id)">Goruotines</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_02_channels" onclick="pathOpen(this.id)">Channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_03_buffered-channels" onclick="pathOpen(this.id)">Buffered channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_04_range-and-close" onclick="pathOpen(this.id)">Channels</a>
                            <a href="/concurrency/4" ng-click="hideTOC(true)" class="ng-binding">Range and Close</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_05_select" onclick="pathOpen(this.id)">Select</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_06_default-selection" onclick="pathOpen(this.id)">Default selection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_07_exercise-equivalent-binary-trees" onclick="pathOpen(this.id)">Exercise: Equivalent binary trees</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_08_mutex-counter" onclick="pathOpen(this.id)">Mutex counter</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a id="home/src/github.com/dougwatson/tour/_07_concurrency/_09_exercise-web-crawler" onclick="pathOpen(this.id)">Exercise: Web crawler</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
    <div class="click-catcher" ng-click="hideTOC(false)"></div>
</div>
