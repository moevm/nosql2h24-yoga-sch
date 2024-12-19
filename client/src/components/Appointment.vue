<<template>
  <div class="page-container">
    <div class="sidebar">
      <h2>Выберите студию</h2>
      <ul class="studio-list">
        <li
            v-for="studio in studios"
            :key="studio.id"
            @click="selectStudio(studio)"
            :class="{ selected: selectedStudio && selectedStudio.id === studio.id }"
        >
          {{ studio.address }}
        </li>
      </ul>
    </div>

    <div class="main-content">
      <template v-if="selectedStudio">
        <div class="calendar-container">
          <h2 id="studio-title">Расписание для {{ selectedStudio.address }}</h2>
          <div class="calendar">
            <div class="calendar-header">
              <button @click="changeMonth(-1)">&#8592;</button>
              <span>{{ months[currentMonth] }} {{ currentYear }}</span>
              <button @click="changeMonth(1)">&#8594;</button>
            </div>
            <div class="calendar-grid">
              <div class="week-days">
                <span>Пн</span>
                <span>Вт</span>
                <span>Ср</span>
                <span>Чт</span>
                <span>Пт</span>
                <span>Сб</span>
                <span>Вс</span>
              </div>
              <div class="dates">
                <div
                    v-for="(date, index) in displayedDates"
                    :key="index"
                    :class="['date', { 'selected': isSelectedDate(date), 'disabled': date.disabled }]"
                    @click="selectDate(date)"
                >
                  {{ date.day }}
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="classes-container">
          <h3 id="classes-date" style="text-align: center">
            Занятия на {{ selectedDate?.day }} {{ months[currentMonth] }} {{ currentYear }}:
          </h3>
          <div class="class-cards">
            <template v-if="classes.length > 0">
              <div v-for="classItem in classes" :key="classItem.id" class="class-card">
                <div class="class-info">
                  <h4>{{ classItem.name }}</h4>
                  <p class="class-time" style="font-size: large">{{ classItem.time }}</p>
                  <p class="class-time">{{ classItem.trainer }}</p>
                </div>
                <button class="enroll-button" @click="enroll(classItem.id)">
                  Записаться
                </button>
              </div>
            </template>
            <template v-else>
              <div class="no-classes">
                Нет занятий на выбранную дату.
              </div>
            </template>
          </div>
        </div>
      </template>

      <template v-else>
        <p>Выберите студию, чтобы увидеть расписание.</p>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';

const URI = `${window.location.protocol}//${window.location.hostname}:${window.location.port}`;

type Studio = {
  id: number;
  address: string;
  classesInfo: { id: number }[];
};

type ClassItem = {
  id: number;
  time: string;
  name: string;
  trainer: string;
};

const studios = ref<Studio[]>([]);
const selectedStudio = ref<Studio | null>(null);
const classes = ref<ClassItem[]>([]);

// Месяцы для отображения
const months = [
  'Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь',
  'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'
];

const currentMonth = ref<number>(new Date().getMonth());
const currentYear = ref<number>(new Date().getFullYear());
const selectedDate = ref<any>(null);

// Даты месяца для отображения
const displayedDates = ref<any[]>([]);

// Функция для получения списка студий
const fetchStudios = async (): Promise<void> => {
  try {
    const response = await fetch(`${URI}/api/v1/studio`);
    if (!response.ok) {
      throw new Error('Ошибка сети');
    }
    let data = await response.json();
    studios.value = data.studios;
  } catch (error) {
    console.error('Ошибка при получении списка студий:', error);
  }
};

const fetchClasses = async (address: string, date: string): Promise<void> => {
  try {
    const selectedDate = new Date(date);

    const timeIntervalBegin = new Date(selectedDate);
    timeIntervalBegin.setHours(0, 0, 0, 0);

    const timeIntervalEnd = new Date(selectedDate);
    timeIntervalEnd.setHours(23, 59, 59, 999);

    const url = `${URI}/api/v1/search/class?studio_address_substrings=${encodeURIComponent(address)}&time_interval_begin=${encodeURIComponent(timeIntervalBegin.toISOString())}&time_interval_end=${encodeURIComponent(timeIntervalEnd.toISOString())}`;

    const response = await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Ошибка сети");
    }

    const data = await response.json();
    console.log(data);
    const selectedClasses = data.classes.map((classData: any) => ({
      id: classData.id,
      time: new Date(classData.time).toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" }),
      name: classData.name,
      trainer: classData.trainerInfo.name,
    }));

    classes.value = selectedClasses.length === 0 ? [] : selectedClasses;
  } catch (error) {
    console.error("Ошибка при получении занятий:", error);
    classes.value = [];
  }
};

const selectStudio = (studio: Studio): void => {
  selectedStudio.value = studio;
  classes.value = [];
  if (selectedDate.value) {
    fetchClasses(studio.address, selectedDate.value.dateString);
  }
};

// Функция для изменения месяца
const changeMonth = (direction: number): void => {
  currentMonth.value += direction;

  if (currentMonth.value < 0) {
    currentMonth.value = 11;
    currentYear.value--;
  } else if (currentMonth.value > 11) {
    currentMonth.value = 0;
    currentYear.value++;
  }

  generateCalendar();
};

// Функция для генерации календаря
const generateCalendar = (): void => {
  const firstDayOfMonth = new Date(currentYear.value, currentMonth.value, 1).getDay();
  const lastDateOfMonth = new Date(currentYear.value, currentMonth.value + 1, 0).getDate();

  const dates = [];
  let day = 1;

  for (let i = 0; i < firstDayOfMonth; i++) {
    dates.push({ day: '', disabled: true });
  }

  for (let i = firstDayOfMonth; i < 7 && day <= lastDateOfMonth; i++) {
    dates.push({ day: day++, disabled: false });
  }

  while (day <= lastDateOfMonth) {
    for (let i = 0; i < 7 && day <= lastDateOfMonth; i++) {
      dates.push({ day: day++, disabled: false });
    }
  }

  displayedDates.value = dates;
};

// Функция для выбора даты
const selectDate = (date: any): void => {
  if (!date.disabled) {
    selectedDate.value = {
      day: date.day,
      dateString: `${currentYear.value}-${String(currentMonth.value + 1).padStart(2, '0')}-${String(date.day).padStart(2, '0')}`
    };
    classes.value = [];
    if (selectedStudio.value) {
      fetchClasses(selectedStudio.value.address, selectedDate.value.dateString);
    }
  }
};

const isSelectedDate = (date: any): boolean => {
  return selectedDate.value && selectedDate.value.dateString === `${currentYear.value}-${String(currentMonth.value + 1).padStart(2, '0')}-${String(date.day).padStart(2, '0')}`;
};

onMounted(() => {
  generateCalendar();
  fetchStudios();
});

watch([selectedStudio, selectedDate], () => {
  if (selectedStudio.value && selectedDate.value) {
    classes.value = [];
    fetchClasses(selectedStudio.value.address, selectedDate.value.dateString);
  }
});

function getCookie(name: string): string | null {
  const matches = document.cookie.match(new RegExp(`(?:^|; )${name.replace(/([.$?*|{}()[\]\\/+^])/g, '\\$1')}=([^;]*)`));
  return matches ? decodeURIComponent(matches[1]) : null;
}

const enroll = async (classId: number): Promise<void> => {
  const cookie = getCookie('User');
  try {
    const response = await fetch(`${URI}/api/v1/appointment`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        clientId: cookie,
        classId: classId,
      }),
    });

    if (!response.ok) {
      throw new Error("Ошибка при записи на занятие");
    }

    alert("Вы успешно записались на занятие!");
  } catch (error) {
    console.error("Ошибка записи:", error);
    alert("Не удалось записаться на занятие. Попробуйте позже.");
  }
};
</script>


<style scoped>
#studio-title {
  text-align: center;
}
.page-container {
  display: flex;
  height: 100vh;
}

.sidebar {
  width: 25%;
  background-color: #f4f4f4;
  padding: 20px;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
}

.studio-list {
  list-style: none;
  padding: 0;
}

.studio-list li {
  padding: 10px;
  cursor: pointer;
  border-bottom: 1px solid #ddd;
}

.studio-list li:hover, .studio-list li.selected {
  background-color: #6A5862;
  color: white;
}

.main-content {
  flex: 1;
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.calendar-container {
  width: 60%;  /* Изменено для ширины календаря */
}

.classes-container {
  width: 35%;  /* Ширина для блока с занятиями */
  padding-left: 20px;
}

.calendar {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #fff;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
}

.calendar-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.calendar-header button {
  padding: 8px 16px;
  font-size: 18px;
  background-color: #e0e0e0;
  color: #6A5862;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.calendar-header button:hover {
  background-color: #cecece;
}

.calendar-grid {
  margin-top: 10px;
  width: 100%;
}

.week-days {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.dates {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 8px;
  margin-top: 15px;
}

.date {
  text-align: center;
  padding: 10px;
  cursor: pointer;
  border-radius: 8px;
  transition: background-color 0.2s ease;
}

.date.selected {
  background-color: #7c6772;
  color: white;
}

.date.disabled {
  color: #ddd;
  cursor: not-allowed;
}

.date:hover {
  background-color: #6A5862;
}

.classes-list {
  list-style: none;
  padding: 0;
}

.classes-list li {
  padding: 10px;
  border-bottom: 1px solid #ddd;
}

.class-cards {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.class-card {
  background: #ffffff;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.enroll-button {
  background: #6A5862;
  color: #ffffff;
  border: none;
  border-radius: 4px;
  padding: 10px 20px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s ease;
}

.enroll-button:hover {
  background: #584b53;
}

.classes-container {
  width: 35%;
  padding-left: 20px;
}

.classes-container.hidden {
  display: none;
}

/* Дополнительные стили */
.class-duration {
  font-size: 12px;
  color: #666;
}

.class-description {
  margin-top: 5px;
  color: #333;
  font-size: 14px;
  line-height: 1.5;
}

.class-instructor {
  font-weight: bold;
  color: #6A5862;
  margin-top: 10px;
}

.filter-bar {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
  background-color: #eaeaea;
  padding: 10px;
  border-radius: 8px;
}

.filter-bar select {
  padding: 8px;
  font-size: 14px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

.filter-bar button {
  padding: 8px 12px;
  background-color: #6A5862;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.filter-bar button:hover {
  background-color: #584b53;
}
</style>
