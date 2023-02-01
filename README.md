<div>
    <ul>
        <li>
            <span>Using the tour</span>
            <ul>
                <li>
                    <span>Welcome!</span>
                    <ul>
                        <li style="list-style-type: none; content: \2713';" >
                            <a class="check" onclick="pathOpen('home/src/github.com/dougwatson/tour/welcome/1/README.md')">Hello, 世界</a>
                        </li>
												<li>
                            <a class="square" onclick="pathOpen('home/src/github.com/dougwatson/tour/welcome/2/README.md')">Go local</a>
                        </li>
												<li>
                            <a class="square" onclick="pathOpen('home/src/github.com/dougwatson/tour/welcome/1/README.md')">Congratulations</a>
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
                        <!-- ngRepeat: p in m.lesson[l].Pages --><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a class="square" href="/basics/1" ng-click="hideTOC(true)" class="ng-binding">Packages</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a class="square" href="/basics/2" ng-click="hideTOC(true)" class="ng-binding">Imports</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/3" ng-click="hideTOC(true)" class="ng-binding">Exported names</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/4" ng-click="hideTOC(true)" class="ng-binding">Functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/5" ng-click="hideTOC(true)" class="ng-binding">Functions continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/6" ng-click="hideTOC(true)" class="ng-binding">Multiple results</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/7" ng-click="hideTOC(true)" class="ng-binding">Named return values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/8" ng-click="hideTOC(true)" class="ng-binding">Variables</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/9" ng-click="hideTOC(true)" class="ng-binding">Variables with initializers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope active" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/10" ng-click="hideTOC(true)" class="ng-binding">Short variable declarations</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/11" ng-click="hideTOC(true)" class="ng-binding">Basic types</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/12" ng-click="hideTOC(true)" class="ng-binding">Zero values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/13" ng-click="hideTOC(true)" class="ng-binding">Type conversions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/14" ng-click="hideTOC(true)" class="ng-binding">Type inference</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/15" ng-click="hideTOC(true)" class="ng-binding">Constants</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/16" ng-click="hideTOC(true)" class="ng-binding">Numeric Constants</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}">
                            <a href="/basics/17" ng-click="hideTOC(true)" class="ng-binding">Congratulations!</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-flowcontrol" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">Flow control statements: for, if, else, switch and defer</span>
                    <ul>
                       <li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/1" ng-click="hideTOC(true)" class="ng-binding">For</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/2" ng-click="hideTOC(true)" class="ng-binding">For continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/3" ng-click="hideTOC(true)" class="ng-binding">For is Go's "while"</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/4" ng-click="hideTOC(true)" class="ng-binding">Forever</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/5" ng-click="hideTOC(true)" class="ng-binding">If</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/6" ng-click="hideTOC(true)" class="ng-binding">If with a short statement</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/7" ng-click="hideTOC(true)" class="ng-binding">If and else</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/8" ng-click="hideTOC(true)" class="ng-binding">Exercise: Loops and Functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/9" ng-click="hideTOC(true)" class="ng-binding">Switch</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/10" ng-click="hideTOC(true)" class="ng-binding">Switch evaluation order</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/11" ng-click="hideTOC(true)" class="ng-binding">Switch with no condition</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/12" ng-click="hideTOC(true)" class="ng-binding">Defer</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/13" ng-click="hideTOC(true)" class="ng-binding">Stacking defers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/flowcontrol/14" ng-click="hideTOC(true)" class="ng-binding">Congratulations!</a>
                        </li>
                    </ul>
                </li>
                <li ng-repeat="l in m.lessons" class="toc-lesson ng-scope" id="toc-l-moretypes" ng-class="{active: l==params.lessonId}">
                    <span ng-click="toggleLesson(l)" class="ng-binding">More types: structs, slices, and maps.</span>
                    <ul>
                        <!-- ngRepeat: p in m.lesson[l].Pages --><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/1" ng-click="hideTOC(true)" class="ng-binding">Pointers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/2" ng-click="hideTOC(true)" class="ng-binding">Structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/3" ng-click="hideTOC(true)" class="ng-binding">Struct Fields</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/4" ng-click="hideTOC(true)" class="ng-binding">Pointers to structs</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/5" ng-click="hideTOC(true)" class="ng-binding">Struct Literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/6" ng-click="hideTOC(true)" class="ng-binding">Arrays</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/7" ng-click="hideTOC(true)" class="ng-binding">Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/8" ng-click="hideTOC(true)" class="ng-binding">Slices are like references to arrays</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/9" ng-click="hideTOC(true)" class="ng-binding">Slice literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/10" ng-click="hideTOC(true)" class="ng-binding">Slice defaults</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/11" ng-click="hideTOC(true)" class="ng-binding">Slice length and capacity</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/12" ng-click="hideTOC(true)" class="ng-binding">Nil slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/13" ng-click="hideTOC(true)" class="ng-binding">Creating a slice with make</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/14" ng-click="hideTOC(true)" class="ng-binding">Slices of slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/15" ng-click="hideTOC(true)" class="ng-binding">Appending to a slice</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/16" ng-click="hideTOC(true)" class="ng-binding">Range</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/17" ng-click="hideTOC(true)" class="ng-binding">Range continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/18" ng-click="hideTOC(true)" class="ng-binding">Exercise: Slices</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/19" ng-click="hideTOC(true)" class="ng-binding">Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/20" ng-click="hideTOC(true)" class="ng-binding">Map literals</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/21" ng-click="hideTOC(true)" class="ng-binding">Map literals continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/22" ng-click="hideTOC(true)" class="ng-binding">Mutating Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/23" ng-click="hideTOC(true)" class="ng-binding">Exercise: Maps</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/24" ng-click="hideTOC(true)" class="ng-binding">Function values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/25" ng-click="hideTOC(true)" class="ng-binding">Function closures</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/26" ng-click="hideTOC(true)" class="ng-binding">Exercise: Fibonacci closure</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/moretypes/27" ng-click="hideTOC(true)" class="ng-binding">Congratulations!</a>
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
                        <!-- ngRepeat: p in m.lesson[l].Pages --><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/1" ng-click="hideTOC(true)" class="ng-binding">Methods</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/2" ng-click="hideTOC(true)" class="ng-binding">Methods are functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/3" ng-click="hideTOC(true)" class="ng-binding">Methods continued</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/4" ng-click="hideTOC(true)" class="ng-binding">Pointer receivers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/5" ng-click="hideTOC(true)" class="ng-binding">Pointers and functions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/6" ng-click="hideTOC(true)" class="ng-binding">Methods and pointer indirection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/7" ng-click="hideTOC(true)" class="ng-binding">Methods and pointer indirection (2)</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/8" ng-click="hideTOC(true)" class="ng-binding">Choosing a value or pointer receiver</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/9" ng-click="hideTOC(true)" class="ng-binding">Interfaces</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/10" ng-click="hideTOC(true)" class="ng-binding">Interfaces are implemented implicitly</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/11" ng-click="hideTOC(true)" class="ng-binding">Interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/12" ng-click="hideTOC(true)" class="ng-binding">Interface values with nil underlying values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/13" ng-click="hideTOC(true)" class="ng-binding">Nil interface values</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/14" ng-click="hideTOC(true)" class="ng-binding">The empty interface</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/15" ng-click="hideTOC(true)" class="ng-binding">Type assertions</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/16" ng-click="hideTOC(true)" class="ng-binding">Type switches</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/17" ng-click="hideTOC(true)" class="ng-binding">Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/18" ng-click="hideTOC(true)" class="ng-binding">Exercise: Stringers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/19" ng-click="hideTOC(true)" class="ng-binding">Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/20" ng-click="hideTOC(true)" class="ng-binding">Exercise: Errors</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/21" ng-click="hideTOC(true)" class="ng-binding">Readers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/22" ng-click="hideTOC(true)" class="ng-binding">Exercise: Readers</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/23" ng-click="hideTOC(true)" class="ng-binding">Exercise: rot13Reader</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/24" ng-click="hideTOC(true)" class="ng-binding">Images</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/25" ng-click="hideTOC(true)" class="ng-binding">Exercise: Images</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/methods/26" ng-click="hideTOC(true)" class="ng-binding">Congratulations!</a>
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
                        <!-- ngRepeat: p in m.lesson[l].Pages --><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/generics/1" ng-click="hideTOC(true)" class="ng-binding">Type parameters</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/generics/2" ng-click="hideTOC(true)" class="ng-binding">Generic types</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/generics/3" ng-click="hideTOC(true)" class="ng-binding">Congratulations!</a>
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
                        <!-- ngRepeat: p in m.lesson[l].Pages --><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/1" ng-click="hideTOC(true)" class="ng-binding">Goroutines</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/2" ng-click="hideTOC(true)" class="ng-binding">Channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/3" ng-click="hideTOC(true)" class="ng-binding">Buffered Channels</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/4" ng-click="hideTOC(true)" class="ng-binding">Range and Close</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/5" ng-click="hideTOC(true)" class="ng-binding">Select</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/6" ng-click="hideTOC(true)" class="ng-binding">Default Selection</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/7" ng-click="hideTOC(true)" class="ng-binding">Exercise: Equivalent Binary Trees</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/8" ng-click="hideTOC(true)" class="ng-binding">Exercise: Equivalent Binary Trees</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/9" ng-click="hideTOC(true)" class="ng-binding">sync.Mutex</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/10" ng-click="hideTOC(true)" class="ng-binding">Exercise: Web Crawler</a>
                        </li><li ng-repeat="p in m.lesson[l].Pages" class="toc-page ng-scope" ng-class="{active: l==params.lessonId &amp;&amp; $index+1==params.pageNumber}" style="display: block;">
                            <a href="/concurrency/11" ng-click="hideTOC(true)" class="ng-binding">Where to Go from here...</a>
                        </li>
                    </ul>
                </li>
            </ul>
        </li>
    </ul>
    <div class="click-catcher" ng-click="hideTOC(false)"></div>
</div>
